FROM golang:1.17 as build
WORKDIR /build
# cache dependencies
ADD go.mod go.sum ./
RUN go mod download
# build
ADD . .
RUN go build -o /build/main

# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2

# provided.al2 has a built-in RIE that runs `exec /var/runtime/boostrap`
# when receiving an event from emulated HTTP endpoint.
# Bootstrap script is usually used for custom runtime such as Rust or Bash,
# but a function built with AWS Lambda Go SDK can be used as such bootstrap
# program. Therefore, simply copying the Go binary to `/var/runtime/boostrap`
# is enough.
COPY --from=build /build/main /var/runtime/bootstrap
CMD [ "/var/runtime/bootstrap" ]

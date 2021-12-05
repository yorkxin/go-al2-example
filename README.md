# Go-Lambda-AL2

Experimental project for running an AWS Lambda function in Go on `provided.al2` runtime.

## Local Testing

### Method 1: Without custom image

Using [Lambda Provided Image](https://gallery.ecr.aws/lambda/provided)

```sh
make lambda
docker run -p 9000:8080 -v "$(pwd)/bin/lambda/bootstrap:/var/runtime/bootstrap" public.ecr.aws/lambda/provided:al2 /var/runime/bootstrap
```

To test event:

```
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"hello":"world"}'
```

### Method 2: With custom image

```sh
docker build -t mytestimg .
docker run -p 9000:8080 mytestimg
```

To test event:

```
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"hello":"world"}'
```

### Method 3: Using SAM

```sh
sam build
sam local start-lambda
```

To test event:

```sh
curl -XPOST "http://localhost:3001/2015-03-31/functions/goal2/invocations" -d '{"hello":"world"}'
```

**NOTE:** URL is different from other methods.

## References

* [Testing Lambda container images locally - AWS Lambda](https://docs.aws.amazon.com/lambda/latest/dg/images-test.html)
* [Deploy Go Lambda functions with container images - AWS Lambda](https://docs.aws.amazon.com/lambda/latest/dg/go-image.html)
* [ECR Public Gallery - AWS Lambda/provided](https://gallery.ecr.aws/lambda/provided)

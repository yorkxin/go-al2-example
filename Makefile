all: lambda

.PHONY: clean
clean:
	rm -rf bin/*

# For running the binary on your machine using provided.al2 image
.PHONY: lambda
lambda:
	GOOS=linux GOARCH=arm64 go build -o bin/lambda/bootstrap main.go

# For directly running `aws lambda update-function-code`
.PHONY: lambda.zip
lambda.zip: lambda
	cd bin/lambda/ && zip lambda.zip bootstrap

# For SAM
.PHONY: build-goal2
build-goal2:
	GOOS=linux go build -o $(ARTIFACTS_DIR)/bootstrap main.go

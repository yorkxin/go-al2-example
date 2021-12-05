all: lambda

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: lambda
lambda:
	GOOS=linux GOARCH=arm64 go build -o bin/lambda.arm64 main.go

.PHONY: lambda.zip
lambda.zip: lambda
	cd bin/ && zip lambda.zip lambda.arm64

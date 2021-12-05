package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type MyEvent struct {
	Hello string `json:"hello"`
}

var stsClient *sts.Client

func NewSTSClient() (*sts.Client, error) {
	if stsClient != nil {
		return stsClient, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration, %v", err)
	}

	stsClient = sts.NewFromConfig(cfg)
	return stsClient, nil
}

func getCallerIdentity() (string, error) {
	client, err := NewSTSClient()

	if err != nil {
		return "", fmt.Errorf("getCallerIdentity: %s", err.Error())
	}

	output, err := client.GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})

	if err != nil {
		return "", fmt.Errorf("getCallerIdentity: %s", err.Error())
	}

	return *output.Arn, nil
}

func handleRequest(ctx context.Context, event MyEvent) (string, error) {
	// event
	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJson)

	// request context
	lc, _ := lambdacontext.FromContext(ctx)
	log.Printf("REQUEST ID: %s", lc.AwsRequestID)
	// global variable
	log.Printf("FUNCTION NAME: %s", lambdacontext.FunctionName)
	// context method
	deadline, _ := ctx.Deadline()
	log.Printf("DEADLINE: %s", deadline)

	arn, err := getCallerIdentity()
	if err != nil {
		return "ERROR", err
	}

	return arn, nil
}

func main() {
	runtime.Start(handleRequest)
}

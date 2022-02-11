package main

import (
	"github.com/aws/aws-lambda-go/events"
	"seed-aws-lambda/handler"
)

func main() {
	handler.NewHandler(events.DynamoDBEvent{})
	handler.NewHandler(events.APIGatewayProxyRequest{})
	handler.NewHandler(events.ALBTargetGroupRequest{})
	handler.NewHandler(events.S3Event{})
}

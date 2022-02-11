package handler

import (
	"reflect"
)

func NewHandler(request interface{}) Handler {
	typeOf := reflect.TypeOf(request)
	switch typeOf.String() {
	case "events.APIGatewayProxyRequest":
		println("hello APIGW")
	case "events.ALBTargetGroupRequest":
		println("hello Target Group")
	case "events.DynamoDBEvent":
		println("hello Dynamo DB")
	case "events.S3Event":
		println("hello s3 event")
	default:
		panic("not implemented")
	}
	return nil
}

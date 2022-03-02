package model

var moduleContent = `module %s 

go 1.17`

var mainContentForLambda = `package main

func main() {
   println("Hello by bava-go, Welcome to project %s in %s")
}

`

var handlerContentLambda = `package handler

import (
		"context"
	"errors"
	"os"
	
	"github.com/aws/aws-lambda-go/events" //Remove import if dont need...
	"github.com/rs/zerolog/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Handler(ctx context.Context, event %s) %s {
	if os.Getenv("ENVIRONMENT") == "local" { //BUG: https://github.com/aws/aws-sam-cli/issues/2510
		ctx = context.TODO()
	}
    span, _ := tracer.StartSpanFromContext(ctx, "<OPERATION_NAME>", tracer.ResourceName("<RESOURCE_NAME>"), //NEED TO EDIT
		tracer.ServiceName("%s-%s"))
	defer span.Finish()
	
	/**
		PUT YOUR CUSTOM TAGS HERE, example 
			span.SetTag("my-tag", myValue)
    */

	logFields := map[string]interface{}{
		"span_id":    span.Context().SpanID(),
		"trace_id":   span.Context().TraceID(),
	}

	logger := log.With().Fields(logFields).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("lambda called...")
	
	//You Logic here...

	logger.Info().Msg("lambda called...Done")

	return errors.New("not implemented"), nil //Remove if dont need
}
`

var awsEventmodel = `package model

type AwsEvent struct {
	Detail struct {
		RequestParameters struct {
			BucketName string 
			Key        string 
		} 
	}
}
`

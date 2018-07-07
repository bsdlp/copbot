package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/bsdlp/copbot/internal/outbound"
	"github.com/bsdlp/copbot/lambdas"
	"github.com/bsdlp/copbot/lambdas/remind/internal"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	cfg := &internal.Config{}
	err := envconfig.Process("remind", cfg)
	if err != nil {
		logger.Fatal(err)
	}

	awsConfig, err := cfg.Config()
	if err != nil {
		logger.Fatal(err)
	}

	remindHandler := &internal.RemindHandler{
		CloudWatchEvents: cloudwatchevents.New(session.New(awsConfig)),
		Outbound: &outbound.Client{
			SQS:      sqs.New(session.New(awsConfig)),
			QueueURL: cfg.OutboundQueueURL,
		},
	}
	lambda.StartHandler(lambdas.MessageHandlerFunc(remindHandler.Create))
}

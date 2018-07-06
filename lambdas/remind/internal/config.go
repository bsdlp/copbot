package internal

import "github.com/bsdlp/copbot/lambdas"

// Config holds configuration options for remind
type Config struct {
	*lambdas.AWSConfig
	OutboundQueueURL string `required:"true"`
}

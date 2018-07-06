package lambdas

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSConfig holds configuration options for aws
type AWSConfig struct {
	Region  string `default:"us-west-2"`
	RoleArn string
}

// Config returns an aws config with the provided options
func (cfg *AWSConfig) Config() (*aws.Config, error) {
	awsConfig := aws.NewConfig().WithRegion(cfg.Region)

	if cfg.RoleArn != "" {
		sess, err := session.NewSession()
		if err != nil {
			return nil, err
		}

		awsConfig = awsConfig.WithCredentials(stscreds.NewCredentials(sess, cfg.RoleArn))
	}
	return awsConfig, nil
}

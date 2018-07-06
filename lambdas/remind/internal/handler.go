package internal

import "github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"

// RemindHandler implements the remind command handler
type RemindHandler struct {
	CloudWatchEvents cloudwatcheventsiface.CloudWatchEventsAPI
}

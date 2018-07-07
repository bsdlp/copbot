package internal

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
	"github.com/bsdlp/copbot/internal/outbound/outboundiface"
)

// RemindHandler implements the remind command handler
type RemindHandler struct {
	CloudWatchEvents cloudwatcheventsiface.CloudWatchEventsAPI
	Outbound         outboundiface.OutboundAPI
}

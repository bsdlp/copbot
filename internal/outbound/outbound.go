package outbound

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/bsdlp/copbot/commands"
	"github.com/golang/protobuf/proto"
)

// Client describes the outbound client
type Client struct {
	SQS      sqsiface.SQSAPI
	QueueURL string
}

// SendMessage implements Outbound
func (c *Client) SendMessage(ctx context.Context, msg *commands.OutboundMessage) (err error) {
	_, err = c.SQS.SendMessageWithContext(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(c.QueueURL),
		MessageBody: aws.String(proto.MarshalTextString(msg)),
	})
	return
}

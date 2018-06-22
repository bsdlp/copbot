package lambdas

import (
	"context"

	"github.com/bsdlp/copbot/commands"
	"github.com/golang/protobuf/proto"
)

// HandlerFunc implements lambda.Handler
type HandlerFunc func(context.Context, *commands.InboundMessage) (*commands.OutboundMessage, error)

// Invoke implements lambda.Handler and handles protobuf marshal/unmarshal
func (hf HandlerFunc) Invoke(ctx context.Context, payload []byte) ([]byte, error) {
	inbound := &commands.InboundMessage{}
	err := proto.Unmarshal(payload, inbound)
	if err != nil {
		return nil, err
	}

	outbound, err := hf(ctx, inbound)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(outbound)
}

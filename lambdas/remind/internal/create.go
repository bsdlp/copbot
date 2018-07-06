package internal

import (
	"context"

	"github.com/bsdlp/copbot/commands"
)

// Create implements the create reminder handler
func (rh *RemindHandler) Create(ctx context.Context, inbound *commands.InboundMessage) (outbound *commands.OutboundMessage, err error) {
	// ListRules
	// if rule already exists:
	//   PutTargets
	// else:
	//   PutRule
	//   PutTargets
	return
}

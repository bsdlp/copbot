package outboundiface

import (
	"context"

	"github.com/bsdlp/copbot/commands"
)

// OutboundAPI describes the outbound client api
type OutboundAPI interface {
	SendMessage(context.Context, *commands.OutboundMessage) error
}

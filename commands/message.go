package commands

import "time"

// DiscordContext contains discord context
type DiscordContext struct {
	DiscordUser      string `json:"discord_user"`
	DiscordMessageID string `json:"discord_message_id"`
	DiscordChannelID string `json:"discord_channel_id"`
}

// ParsedMessage contains the parsed message
type ParsedMessage struct {
	Command   string   `json:"command"`
	Arguments []string `json:"arguments"`
}

// Message contains the message
type Message struct {
	MessageTimestamp time.Time       `json:"message_timestamp"`
	RawMessage       string          `json:"raw_message"`
	Message          *ParsedMessage  `json:"message,omitempty"`
	DiscordContext   *DiscordContext `json:"discord_context,omitempty"`
}

package notify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Notify_DiscordBody_WithBody(t *testing.T) {
	// check new body is set
	{
		data := []string{"dog", "mom", "cat", "act"}
		for _, bod := range data {
			var b DiscordBody
			b = b.WithBody([]byte(bod))
			assert.Equal(t, bod, b.Content)
		}
	}
}

func Test_Notify_NewDiscordSender(t *testing.T) {
	// check we get back a non nil ptr and the configs make sense
	{
		cfg := &SlackConfig{
			DefaultChannelWebhook: "some default channel",
		}
		sender := NewDiscordSender(cfg)
		assert.NotNil(t, sender)
		assert.NotNil(t, sender.r)
		assert.Equal(t, discordSenderErrPrefix, sender.errPrefix)
		assert.Equal(t, sender.defaultWebhookURL, cfg.DefaultChannelWebhook)
	}
}

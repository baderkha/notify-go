package notify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Notify_SlackBody_WithBody(t *testing.T) {
	// check new body is set
	{
		data := []string{"dog", "mom", "cat", "act"}
		for _, bod := range data {
			var b SlackBody
			b = b.WithBody([]byte(bod))
			assert.Equal(t, bod, b.Text)
		}
	}
}

func Test_Notify_NewSlackSender(t *testing.T) {
	// check we get back a non nil ptr and the configs make sense
	{
		cfg := &SlackConfig{
			DefaultChannelWebhook: "some default channel",
		}
		sender := NewSlackSender(cfg)
		assert.NotNil(t, sender)
		assert.NotNil(t, sender.r)
		assert.Equal(t, slackSenderErrPrefix, sender.errPrefix)
		assert.Equal(t, sender.defaultWebhookURL, cfg.DefaultChannelWebhook)
	}
}

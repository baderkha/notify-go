package notify

var _ MessageSender = &SlackSender{}

const discordSenderErrPrefix = "notify-go : discord message error : "

// NewDiscordSender : creat a new discord notification sender
func NewDiscordSender(cfg *DiscordConfig) *DiscordSender {
	return (NewWebhookSender[DiscordBody](cfg, discordSenderErrPrefix))
}

// DiscordConfig : Discord configuration
type DiscordConfig = WebhookConfig

// DiscordBody : Discord body
type DiscordBody struct {
	Content string `json:"content"`
}

func (s DiscordBody) WithBody(b []byte) DiscordBody {
	s.Content = string(b)
	return s
}

// DiscordSender : slack message sender
type DiscordSender = WebhookSender[DiscordBody]

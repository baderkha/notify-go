package notify

var _ MessageSender = &SlackSender{}

const discordSenderErrPrefix = "notify-go : discord message error : "

// NewDiscordSender : creat a new discord notification sender
func NewDiscordSender() *DiscordSender {
	return (NewWebhookSender[DiscordBody](discordSenderErrPrefix))
}

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

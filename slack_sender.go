package notify

var _ MessageSender = &SlackSender{}

const slackSenderErrPrefix = "notify-go : slack message error : "

// NewSlackSender : creat a new slack notification sender
func NewSlackSender(cfg *SlackConfig) *SlackSender {
	return (NewWebhookSender[SlackBody](cfg, slackSenderErrPrefix))
}

// SlackConfig : slack configuration
type SlackConfig = WebhookConfig

type SlackBody struct {
	Text string `json:"text"`
}

func (s SlackBody) WithBody(b []byte) SlackBody {
	s.Text = string(b)
	return s
}

// SlackSender : slack message sender
type SlackSender = WebhookSender[SlackBody]

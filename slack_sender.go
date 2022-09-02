package notify

var _ MessageSender = &SlackSender{}

const slackSenderErrPrefix = "notify-go : slack message error : "

// NewSlackSender : creat a new slack notification sender
func NewSlackSender() *SlackSender {
	return (NewWebhookSender[SlackBody](slackSenderErrPrefix))
}

type SlackBody struct {
	Text string `json:"text"`
}

func (s SlackBody) WithBody(b []byte) SlackBody {
	s.Text = string(b)
	return s
}

// SlackSender : slack message sender
type SlackSender = WebhookSender[SlackBody]

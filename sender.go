package notify

import "os"

var (
	debugging = os.Getenv(EnvDebugging) == "1" || os.Getenv(EnvDebugging) == "TRUE"
)

func SetDebug(val bool) {
	debugging = val
}

type WebhookConfig struct {
	DefaultChannelWebhook string
}

// MessageSender : a base contract that all other message
// senders need to fullfill
type MessageSender interface {
	// SendToDefaultReciever : send to someone in your default configs
	SendToDefaultReciever(bodyContent []byte) error
	// SendToReciever send to someone not in your default configs
	SendToReciever(reciever string, bodyContent []byte) error
}

type MessageManager struct {
	senders map[string]MessageSender
}

func (m *MessageManager) SendDefaultAll(body []byte) error

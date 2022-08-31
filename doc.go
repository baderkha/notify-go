// Package notify: a notificaion publishing package that allows you to send messages to different social platforms
package notify

const (

	// EnvDebugging : setting this environment
	// enables debugging logs to stdout
	EnvDebugging = "NOTIFY_GO_DEBUGGING"
)

const (
	DiscordType  = "discord"
	SlackType    = "slack"
	TelegramType = "telegram"
)

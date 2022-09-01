// Package notify: a notificaion publishing package that allows you to send messages to different social platforms
package notify

import "strings"

const (

	// EnvHTTPDebugging : setting this environment
	// enables debugging http logs to stdout
	EnvHTTPDebugging = "NOTIFY_GO_HTTP_DEBUGGING"
)

const (
	// DiscordType : senderType of discord
	DiscordSenderType = "discord"
	// SlackSenderType : senderType of Slack
	SlackSenderType = "slack"
	// TelegramSenderType : senderType of telegram
	TelegramSenderType = "telegram"
	// TestSenderType : sender Type for testing
	TestSenderType = "test"
)

var (
	SupportedTypes = []string{
		DiscordSenderType,
		SlackSenderType,
		TelegramSenderType,
		TestSenderType,
	}
	SupportedTypesString = strings.Join(SupportedTypes, " , ")
)

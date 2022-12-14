package notify

import (
	"os"
)

var (
	isDebug = os.Getenv(EnvHTTPDebugging) == "1" || os.Getenv(EnvHTTPDebugging) == "TRUE"
)

// SetDebug : turn debug on or off
//
// Otherwise it will default to the environment variable
// see : EnvHTTPDebugging for the env var key
func SetDebug(val bool) {
	isDebug = val
}

// MessageSender : a base contract that all other message
// senders need to fullfill
type MessageSender interface {
	// Send send to someone not in your default configs
	Send(reciever string, bodyContent []byte) error
}

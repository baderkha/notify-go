package notify

import (
	"fmt"
	"os"
	"sync"

	"github.com/baderkha/notify-go/pkg/thirdparty/errors"
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

func NewRecieverAlias(senderTypeToReciverMap map[string]string) *RecieverAlias {
	return &RecieverAlias{
		aliasMap: senderTypeToReciverMap,
	}
}

type RecieverAlias struct {
	aliasMap map[string]string
}

func (r *RecieverAlias) get(senderType string) string {
	return r.aliasMap[senderType]
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

// SendDefaultAll : concurrently sends messages to all the setup senders
func (m *MessageManager) SendDefaultAll(body []byte) error {
	var wg sync.WaitGroup
	var erList errors.ErrorList
	for _, v := range m.senders {
		wg.Add(1)
		go func(body []byte) {
			defer wg.Done()
			err := v.SendToDefaultReciever(body)
			if err != nil {
				erList.Push(err)
			}
		}(body)
	}
	wg.Wait()
	if erList.Len() > 0 {
		return erList.Err()
	}
	return nil
}

// SendAllToSameReciever : concurrently sends messages to all the senders given a reciver alias address map
func (m *MessageManager) SendAllToSameReciever(alias *RecieverAlias, bodyContent []byte) error {
	var wg sync.WaitGroup
	var erList errors.ErrorList
	for k, v := range m.senders {
		wg.Add(1)
		go func(body []byte, alias *RecieverAlias, senderType string, v MessageSender) {
			defer wg.Done()
			reciever := alias.get(senderType)
			if reciever == "" {
				erList.Push(fmt.Errorf("notfy-go : Message Manager : Expected alias entry for senderType %s", senderType))
				return
			}
			err := v.SendToReciever(reciever, body)
			if err != nil {
				erList.Push(err)
			}
		}(bodyContent, alias, k, v)
	}
	wg.Wait()
	if erList.Len() > 0 {
		return erList.Err()
	}
	return nil
}

// SendToSpecificSenderType : allows you to access the SendToReciever Method
func (m *MessageManager) SendToSpecificSenderType(senderType, reciever string, bodyContent []byte) error {
	sender := m.senders[senderType]
	if sender == nil {
		return fmt.Errorf("notify-go : Message Manager : Expected %s to be setup", senderType)
	}
	return sender.SendToReciever(reciever, bodyContent)
}

// SendToSpecificSenderTypeDefault : Send to a specific social type ie discord forexample , using the default reciver
// ie the channel / group / ...etc
func (m *MessageManager) SendToSpecificSenderTypeDefault(senderType string, bodyContent []byte) error {
	sender := m.senders[senderType]
	if sender == nil {
		return fmt.Errorf("notify-go : Message Manager : Expected %s to be setup", senderType)
	}
	return sender.SendToDefaultReciever(bodyContent)
}

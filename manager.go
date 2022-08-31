package notify

import (
	"fmt"
	"sync"

	"github.com/baderkha/notify-go/pkg/thirdparty/errors"
)

// Manager : A manager for all the sender type implementations
type Manager struct {
	senders map[string]MessageSender
}

// SendDefaultAll : concurrently sends messages to all the setup senders
func (m *Manager) SendDefaultAll(body []byte) error {
	var wg sync.WaitGroup
	var erList errors.ErrorList
	for _, v := range m.senders {
		wg.Add(1)
		go func(body []byte, sender MessageSender) {
			defer wg.Done()
			err := sender.SendToDefaultReciever(body)
			if err != nil {
				erList.Push(err)
			}
		}(body, v)
	}
	wg.Wait()
	if erList.Len() > 0 {
		return erList.Err()
	}
	return nil
}

// SendAllToSameReciever : concurrently sends messages to all the senders given a reciver alias address map
func (m *Manager) SendAllToSameReciever(alias *RecieverAlias, bodyContent []byte) error {
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
func (m *Manager) SendToSpecificSenderType(senderType, reciever string, bodyContent []byte) error {
	sender := m.senders[senderType]
	if sender == nil {
		return fmt.Errorf("notify-go : Message Manager : Expected %s to be setup", senderType)
	}
	return sender.SendToReciever(reciever, bodyContent)
}

// SendToSpecificSenderTypeDefault : Send to a specific social type ie discord forexample , using the default reciver
// ie the channel / group / ...etc
func (m *Manager) SendToSpecificSenderTypeDefault(senderType string, bodyContent []byte) error {
	sender := m.senders[senderType]
	if sender == nil {
		return fmt.Errorf("notify-go : Message Manager : Expected %s to be setup", senderType)
	}
	return sender.SendToDefaultReciever(bodyContent)
}

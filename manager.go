package notify

import (
	"errors"
	"fmt"
	"sync"

	err "github.com/baderkha/notify-go/pkg/thirdparty/err"
)

var (
	errExpectedAlias = errors.New("notify-go : Error expected an alias mapping to be provided")
)

// Manager : A manager for all the sender type implementations
type Manager struct {
	senders map[string]MessageSender
	mu      sync.Mutex
}

// AddSender : add a sender service
// Try to use this only during your app's init phase
// preferably
func (m *Manager) AddSender(senderType string, s MessageSender) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.senders == nil {
		m.senders = make(map[string]MessageSender)
	}
	m.senders[senderType] = s
}

// SendAll : concurrently sends messages to all the senders given a reciver alias address map
func (m *Manager) SendAll(alias *RecieverAlias, bodyContent []byte) error {
	if alias == nil {
		return errExpectedAlias
	}
	var wg sync.WaitGroup
	var erList err.List
	for k, v := range m.senders {
		wg.Add(1)
		go func(body []byte, alias *RecieverAlias, senderType string, v MessageSender) {
			defer wg.Done()
			reciever := alias.Get(senderType)
			if reciever == "" {
				erList.Push(fmt.Errorf("notfy-go : Message Manager : Expected alias entry for senderType %s", senderType))
				return
			}
			err := v.Send(reciever, body)
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

// SendToSpecificSenderType : allows you to access the Send Method
func (m *Manager) SendToSpecificSenderType(senderType, reciever string, bodyContent []byte) error {
	sender := m.senders[senderType]
	if sender == nil {
		return fmt.Errorf("notify-go : Message Manager : Expected %s to be setup", senderType)
	}
	return sender.Send(reciever, bodyContent)
}

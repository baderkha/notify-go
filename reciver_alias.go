package notify

import (
	"fmt"
	"strings"
	"sync"
)

// NewRecieverAlias : returns a reciver alias
// A reciever alias is basically a map of sender type -> reciever
func NewRecieverAlias(senderTypeToReciverMap map[string]string) (*RecieverAlias, error) {
	for k, v := range senderTypeToReciverMap {
		if !strings.Contains(SupportedTypesString, k) {
			return nil, fmt.Errorf("notify-go : error this sender type %s is not supported at the moment", k)
		}

		if v == "" {
			return nil, fmt.Errorf("notify-go : error this sender type %s does not have a reciver", k)
		}
	}
	return &RecieverAlias{
		aliasMap: senderTypeToReciverMap,
	}, nil
}

func NewEmptyRecieverAlias() *RecieverAlias {
	return &RecieverAlias{
		aliasMap: make(map[string]string),
	}
}

// RecieverAlias : a map of a sender and reciver, it's a good way to group similar recivers but different channels
//
// In plain english : let's say you had the same group in discord , slack , telegram
// this struct will provide you with functionality to be able to map the same group for different channels
//
// It's go routine safe !
//
//	forexample :
//
//	// from primitive data
//	var s map[string]string = make(map[string]string)
//	s[notify.DiscordSenderType] = "https://somechannelwebhook.com"
//	s[notify.TestSenderType] = "some_channel_id"
//	// will throw an error if sender type is not supported or if reciver is empty for sender
//	coolCryptoChannel,err := notify.NewReciverAlias(s)
//
//	// using the ReciverAlias struct (it's go routine safe)
//	// getting reciver
//	DiscordChannelWebhook := coolCryptoChannel.Get(notify.DiscordSenderType)
//
//	// enrolling new reciver
//	coolCryptoChannel.Add(notify.SlackSenderType,"https://someslackchannelwebooh.com")
type RecieverAlias struct {
	aliasMap map[string]string
	mu       sync.RWMutex
}

// Get : get a reciver for a sender type
func (r *RecieverAlias) Get(senderType string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	res := r.aliasMap[senderType]
	return res
}

// Add : add a sender -> reciver mapping
// will throw error if sender type is not supported / reciver is empty
func (r *RecieverAlias) Add(senderType string, reciver string) error {
	if !strings.Contains(SupportedTypesString, senderType) {
		return fmt.Errorf("notify-go : error this sender type %s is not supported at the moment", senderType)
	}
	if reciver == "" {
		return fmt.Errorf("notify-go : error this sender type %s does not have a reciver", senderType)
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.aliasMap == nil {
		r.aliasMap = make(map[string]string)
	}
	r.aliasMap[senderType] = reciver
	return nil
}

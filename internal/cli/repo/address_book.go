package repo

// Address : address model
type Address struct {
	Name    string            `json:"name"`    // a user friendly name for this channel ie my_crypto_channel_in_slack_and_discord
	Socials map[string]string `json:"socials"` // mapping of sender -> reciever ie slack=>https://webhook.com
}

// IAddressBook : address book repo
type IAddressBook interface {
	GetEntireAddressBook() []*Address
	GetByLabel(label string) (a *Address, isFound bool)
	Add(m *Address)                  // just panic if can't
	Update(label string, m *Address) // just panic if can't
	RemoveByLabel(label string)      // just panic if can't
	DeleteAll()                      // just panic if can't
	// after doing transactions , hit the flush button to you know
	// flush
	Flush()
}

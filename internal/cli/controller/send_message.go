package controller

import (
	"fmt"
	"sync"

	"github.com/baderkha/notify-go"
	"github.com/baderkha/notify-go/internal/cli/repo"
	"github.com/spf13/cobra"
)

var (
	notifyMgr *notify.Manager
)

func SendMessage(cmd *cobra.Command, args []string) {
	senderType := args[0]
	contactNameOrWH := args[1]
	messageContent := args[2]

	a, isFound := contactRepo.GetByLabel(contactNameOrWH)

	if isFound && a.Socials[senderType] != "" {
		contactNameOrWH = a.Socials[senderType]
	}
	err := notifyMgr.SendToSpecificType(senderType, contactNameOrWH, []byte(messageContent))
	if err != nil {
		cmd.PrintErr(err)
		fmt.Println("")
	}
	cmd.Println("sent message !")
}

func SendMessageToAllChannels(cmd *cobra.Command, args []string) {
	contactName := args[0]
	messageContent := args[1]

	a, isFound := contactRepo.GetByLabel(contactName)

	if !isFound {
		cmd.PrintErr(fmt.Sprintf("contact '%s' not found \n", contactName))
		return
	}

	ralias, err := notify.NewRecieverAlias(a.Socials)
	if err != nil {
		panic(err) // not supposed to happen
	}

	err = notifyMgr.SendAll(ralias, []byte(messageContent))
	if err != nil {
		cmd.PrintErr(err.Error() + "\n")
		return
	}

}

func BroadCastMessageToAll(cmd *cobra.Command, args []string) {
	messageContent := args[0]

	var wg sync.WaitGroup

	a := contactRepo.GetEntireAddressBook()

	for _, contact := range a {
		if contact == nil {
			continue
		}
		wg.Add(1)
		go func(con *repo.Address) {
			defer wg.Done()
			ralias, err := notify.NewRecieverAlias(con.Socials)
			if err != nil {
				panic(err) // not supposed to happen
			}

			err = notifyMgr.SendAll(ralias, []byte(messageContent))
			if err != nil {
				cmd.PrintErr(err.Error() + "\n")
				return
			}
		}(contact)

	}
	wg.Wait()

}

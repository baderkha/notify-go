package controller

import (
	"fmt"

	"github.com/baderkha/notify-go"
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
	err := notifyMgr.SendToSpecificSenderType(senderType, contactNameOrWH, []byte(messageContent))
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

package controller

import (
	"fmt"

	"github.com/baderkha/notify-go/internal/cli/repo"
	"github.com/chrusty/go-tableprinter"
	"github.com/spf13/cobra"
)

var (
	contactRepo repo.IAddressBook
)

func NewContact(cmd *cobra.Command, args []string) {
	contactName := args[0]
	_, isFound := contactRepo.GetByLabel(contactName)
	if isFound {
		cmd.Print("This Contact already exists \n")
		return
	}
	contactRepo.Add(&repo.Address{
		Name:    contactName,
		Socials: make(map[string]string),
	})
	contactRepo.Flush()
}

func AppendContact(cmd *cobra.Command, args []string) {
	contactName := args[0]
	socialType := args[1]
	value := args[2]
	contactSocials, isFound := contactRepo.GetByLabel(contactName)
	if !isFound {
		cmd.PrintErr("You need to create the contact first \n")
		return
	}
	if contactSocials.Socials == nil {
		contactSocials.Socials = make(map[string]string)
	}
	contactSocials.Socials[socialType] = value
	contactRepo.Update(contactName, &repo.Address{
		Name:    contactName,
		Socials: contactSocials.Socials,
	})
	contactRepo.Flush()
}

func AllContact(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		ad, isFound := contactRepo.GetByLabel(args[0])

		if !isFound {
			fmt.Println("Contact not found")
			return
		}
		fmt.Println("")
		tableprinter.Print(ad)
		fmt.Println("")
		return
	}
	ad := contactRepo.GetEntireAddressBook()
	if ad == nil {
		ad = make([]*repo.Address, 0)
	}
	fmt.Println("")
	tableprinter.Print(ad)
	fmt.Println("")
}

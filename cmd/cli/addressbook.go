package cli

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/baderkha/notify-go"
	"github.com/baderkha/notify-go/internal/cli/controller"
	"github.com/spf13/cobra"
)

var (
	errAppendArgs         = errors.New("append Contact needs a friendly name , social type , and a chatid/webhook value")
	errNewCon             = errors.New("new-contact needs a name for the contact")
	errBadURLSlackDiscord = errors.New("slack and discord require a valid webhook url")
)

func isSupportedType(socialType string) error {
	if strings.Contains(notify.SupportedTypesString, socialType) {
		return nil
	}
	return fmt.Errorf("social type must be either %s", notify.SupportedTypesString)
}

func isSupportedValue(socialType, socialTypeValue string) error {
	switch socialType {
	case notify.DiscordSenderType:
		fallthrough
	case notify.SlackSenderType:
		_, err := url.ParseRequestURI(socialTypeValue)
		if err != nil {
			return errBadURLSlackDiscord
		}
		return nil
	default:
		return nil
	}
}

var (
	appendContact = &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 3 {
				return errAppendArgs
			}
			err := isSupportedType(args[1])
			if err != nil {
				return err
			}
			err = isSupportedValue(args[1], args[2])
			if err != nil {
				return err
			}
			return nil
		},
		Use:     fmt.Sprintf("append-contact-social [<<contact name>>] [%s] [chatid|webhookurl]", strings.Join(notify.SupportedTypes, "|")),
		Aliases: []string{"apcon"},
		Short:   "append an address binding for a social platform",
		Long:    "append a social webhook / chat id for a friendly name you created",
		Run:     controller.AppendContact,
	}

	addContact = &cobra.Command{
		Use: "new-contact [<<contact name>>]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errNewCon
			}
			return nil
		},
		Aliases: []string{"newcon"},
		Short:   "Add a contact ",
		Long:    "a contact is basically a user friendly name => a map of social types => webhook values/chatId values",
		Run:     controller.NewContact,
	}

	allContact = &cobra.Command{
		Use:     "contacts",
		Aliases: []string{"cons"},
		Short:   "prints all the contacts you have",
		Run:     controller.AllContact,
	}
)

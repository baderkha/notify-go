package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/baderkha/notify-go"
	"github.com/baderkha/notify-go/internal/cli/controller"
	"github.com/spf13/cobra"
)

var (
	errSendMsgArgs      = errors.New("needs to have sender type , contact name | chat id | webhook url , message body")
	errSendMsgToConArgs = errors.New("needs to have contact name | chat id | webhook url , message body")
	errSendMsgToAllArgs = errors.New("needs to have a message body")
)

var (
	sendMessage = &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 3 {
				return errSendMsgArgs
			}
			err := isSupportedType(args[0])
			if err != nil {
				return err
			}
			return nil
		},
		Use:     fmt.Sprintf("send-message  [%s] [<<contact name>> |<<chatid>>|<<webhook url>>] [message body]", strings.Join(notify.SupportedTypes, "|")),
		Aliases: []string{"msg"},
		Short:   "send a message with discord , slack or telegram",
		Run:     controller.SendMessage,
	}

	sendMessageToContact = &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errSendMsgToConArgs
			}
			return nil
		},
		Use:     "send-message-to-con [<<contact name>> |<<chatid>>|<<webhook url>>] [message body]",
		Aliases: []string{"msgcon"},
		Short:   "send a message with discord , slack or telegram  to a specified contact",
		Run:     controller.SendMessageToAllChannels,
	}

	broadcastToAll = &cobra.Command{
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errSendMsgToAllArgs
			}
			return nil
		},
		Use:     "broadcast-message-to-all [message body]",
		Aliases: []string{"msgbrod"},
		Short:   "send a message with discord , slack or telegram for you entire address book",
		Run:     controller.BroadCastMessageToAll,
	}
)

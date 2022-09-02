package cli

import (
	"fmt"
	"strings"

	"github.com/baderkha/notify-go"
	"github.com/baderkha/notify-go/internal/cli/controller"
	"github.com/spf13/cobra"
)

var (
	supportedTypesCompact = strings.Join(notify.SupportedTypes, " , ")
	prettySupportedTypes  = strings.Join(notify.SupportedTypes, "\n - ")

	rootCmd = &cobra.Command{
		Use:   "notify-go",
		Short: fmt.Sprintf("A bot-client message-sender for %s ", supportedTypesCompact),
		Long: fmt.Sprintf(`

notify-go is a bot-client sender that allows you to 
orchestrate sending messages to many supported social platforms.

Supported Platforms are as follows : 
 - %s

You can choose to send to all clients or a specific one. 

Setup is done once , and after that you can automate your message sending


		`, prettySupportedTypes),
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	controller.Init()
	rootCmd.AddCommand(appendContact)
	rootCmd.AddCommand(addContact)
	rootCmd.AddCommand(allContact)

	rootCmd.AddCommand(sendMessage)
	rootCmd.AddCommand(sendMessageToContact)
	rootCmd.AddCommand(broadcastToAll)
}

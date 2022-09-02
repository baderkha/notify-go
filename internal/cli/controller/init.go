package controller

import (
	"github.com/baderkha/notify-go"
	"github.com/baderkha/notify-go/internal/cli/pkg/config"
	"github.com/baderkha/notify-go/internal/cli/repo"
	"github.com/baderkha/notify-go/pkg/serializer"
)

func Init() {

	contactRepo = (&repo.AddressBookFile{
		Slizr:     serializer.JSON[[]*repo.Address]{},
		WritePath: config.GetPath(),
	}).Init()

	notifyMgr = new(notify.Manager)
	notifyMgr.AddSender(notify.DiscordSenderType, notify.NewDiscordSender())
	notifyMgr.AddSender(notify.SlackSenderType, notify.NewSlackSender())
}

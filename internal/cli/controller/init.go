package controller

import (
	"github.com/baderkha/notify-go"
	"github.com/baderkha/notify-go/internal/cli/pkg/config"
	"github.com/baderkha/notify-go/internal/cli/repo"
	"github.com/baderkha/notify-go/pkg/serializer"
)

func Init() {
	config.InitFolderPath()
	contactRepo = (&repo.AddressBookFile{
		Slizr:     serializer.JSON[[]*repo.Address]{},
		WritePath: config.GetPath(),
	}).Init()

	notifyMgr = notify.Default()
}

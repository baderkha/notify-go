package controller

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/baderkha/notify-go/internal/cli/repo"
	"github.com/baderkha/notify-go/pkg/serializer"
	"github.com/chrusty/go-tableprinter"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func addressBook() *repo.AddressBookFile {
	tmpPath := fmt.Sprintf("/tmp/notify-go/testing/%s/", uuid.New().String())
	os.MkdirAll(tmpPath, os.ModePerm)
	contactRepo = &repo.AddressBookFile{
		Slizr:     serializer.JSON[[]*repo.Address]{},
		WritePath: tmpPath,
	}
	return contactRepo.(*repo.AddressBookFile)
}

func initAddress(rpo *repo.AddressBookFile, name string) {
	rpo.Add(&repo.Address{
		Name: name,
	})
	rpo.Flush()
}

func Test_Controller_NewContact(t *testing.T) {
	dummyName := "dummy_name"
	// if contact already exists , EXPECT NO CALL TO THE NEW REPO METHOD
	{
		rpo := addressBook().Init()
		rpo.Add(&repo.Address{
			Name: dummyName,
		})
		rpo.Flush()
		NewContact(&cobra.Command{}, []string{dummyName})

		assert.Len(t, rpo.GetEntireAddressBook(), 1)
	}
	// if contact doesn't exist , expect a new address name with empty map
	{
		rpo := addressBook().Init()
		NewContact(&cobra.Command{}, []string{dummyName})

		assert.Len(t, rpo.GetEntireAddressBook(), 1)
		assert.Equal(t, dummyName, rpo.GetEntireAddressBook()[0].Name)
		assert.Len(t, rpo.GetEntireAddressBook()[0].Socials, 0)
	}
}

func Test_Controller_AllContact(t *testing.T) {
	dummyName := "dummy_name"
	dummyName2 := "dummy_name_2"

	// if not specific then get back all , expect table printer to print same info
	{
		// if data is there
		{
			outExpected := new(bytes.Buffer)
			outActual := new(bytes.Buffer)
			rpo := addressBook().Init()
			initAddress(rpo, dummyName)
			initAddress(rpo, dummyName2)
			tableprinter.SetOutput(outExpected)
			tableprinter.Print(rpo.GetEntireAddressBook())
			tableprinter.SetOutput(outActual)
			AllContact(&cobra.Command{}, nil)

			assert.Equal(t, outExpected.String(), outActual.String())
		}
		// if empty expect empty array
		{
			outExpected := new(bytes.Buffer)
			outActual := new(bytes.Buffer)
			addressBook().Init()
			tableprinter.SetOutput(outExpected)
			tableprinter.Print(make([]*repo.Address, 0))
			tableprinter.SetOutput(outActual)
			AllContact(&cobra.Command{}, nil)

			assert.Equal(t, outExpected.String(), outActual.String())
		}
	}
	// if requested a specific person , and not found  no table printed
	{
		outExpected := new(bytes.Buffer)
		outActual := new(bytes.Buffer)
		rpo := addressBook().Init()
		initAddress(rpo, dummyName)
		initAddress(rpo, dummyName2)
		tableprinter.SetOutput(outExpected)
		tableprinter.Print(nil)
		tableprinter.SetOutput(outActual)
		AllContact(&cobra.Command{}, []string{"some_unknown_name"})

		assert.Equal(t, outExpected.String(), outActual.String())
	}

	// if requested a specific person , and found , then print table for that person
	{
		outExpected := new(bytes.Buffer)
		outActual := new(bytes.Buffer)
		rpo := addressBook().Init()
		initAddress(rpo, dummyName)
		initAddress(rpo, dummyName2)
		tableprinter.SetOutput(outExpected)
		info, _ := rpo.GetByLabel(dummyName)
		tableprinter.Print(info)
		tableprinter.SetOutput(outActual)
		AllContact(&cobra.Command{}, []string{dummyName})
		assert.Equal(t, outExpected.String(), outActual.String())
	}
}

func Test_Controller_AppendContact(t *testing.T) {

}

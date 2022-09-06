package repo

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/baderkha/notify-go/pkg/serializer"
	"github.com/davecgh/go-spew/spew"
)

var _ IAddressBook = &AddressBookFile{}

const addBookFileName = ".address_book.json"

type AddressBookFile struct {
	Slizr     serializer.JSON[[]*Address]
	WritePath string
	books     []*Address
	mu        sync.RWMutex
}

func (a *AddressBookFile) path() string {
	addressPath := filepath.Join(a.WritePath, addBookFileName)
	return addressPath
}

func (a *AddressBookFile) Init() *AddressBookFile {
	a.mu.Lock()
	defer a.mu.Unlock()
	f, _ := os.Open(a.path())
	defer f.Close()
	add, err := a.Slizr.Read(f)
	if err != nil || add == nil {
		add = &[]*Address{}

		os.Remove(a.path())

		fw, err := os.Create(a.path())
		spew.Dump(err)
		defer fw.Close()
		err = a.Slizr.Write(add, fw)
		if err != nil {
			panic(err)
		}
	}

	a.books = *add
	return a
}

func (a *AddressBookFile) musteWriteAddressBook(ad []*Address) {
	f, _ := os.OpenFile(a.path(), os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, os.ModePerm)
	f.Truncate(0)
	f.Seek(0, 0)
	defer f.Close()
	err := a.Slizr.Write(&ad, f)
	if err != nil {
		panic(err)
	}
	a.books = ad
}

func (a *AddressBookFile) Flush() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.musteWriteAddressBook(a.books)
}

func (a *AddressBookFile) GetEntireAddressBook() []*Address {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.books
}
func (a *AddressBookFile) GetByLabel(label string) (ad *Address, isFound bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, adrs := range a.books {
		if adrs.Name == label {
			return adrs, true
		}
	}
	return nil, false
}
func (a *AddressBookFile) Add(m *Address) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.books = append(a.books, m)

}
func (a *AddressBookFile) Update(label string, m *Address) {
	a.mu.Lock()
	defer a.mu.Unlock()
	addresses := a.books
	for i, adrs := range addresses {
		if adrs.Name == label {
			addresses[i] = m
		}
	}
	a.books = addresses

}
func (a *AddressBookFile) RemoveByLabel(label string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	newAdrs := []*Address{}
	for _, adrs := range a.books {
		if adrs.Name == label {
			continue
		}
		newAdrs = append(newAdrs, adrs)
	}
	a.books = newAdrs
}

func (a *AddressBookFile) DeleteAll() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.books = []*Address{}
}

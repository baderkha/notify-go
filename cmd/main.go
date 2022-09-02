package main

import "github.com/baderkha/notify-go/cmd/cli"

func main() {
	cli.Execute()
	//cli.Execute()

	// config.InitFolderPath()
	// adBook := repo.AddressBookFile{
	// 	Slizr:     serializer.JSON[[]*repo.Address]{},
	// 	WritePath: config.GetPath(),
	// }
	// var mu sync.WaitGroup
	// adBook.Init()
	// start := time.Now()
	// for i := 0; i < 100; i++ {
	// 	mu.Add(1)
	// 	go func(i int) {

	// 		defer mu.Done()
	// 		fmt.Println("started Routine => " + fmt.Sprint(i))
	// 		// adBook.Update("ahmad", &repo.Address{
	// 		// 	Label:     "ahmad",
	// 		// 	SenderMap: make(map[string]string),
	// 		// })

	// 		// adBook.GetByLabel("ahmad")
	// 		if i%2 == 0 {
	// 			adBook.Add(&repo.Address{
	// 				Label:     "ahmad",
	// 				SenderMap: make(map[string]string),
	// 			})
	// 		} else {
	// 			adBook.GetByLabel("ahmad")
	// 		}

	// 	}(i)
	// }
	// mu.Wait()
	// adBook.Flush()
	// elapsed := time.Since(start)
	// spew.Dump(elapsed)
	// spew.Dump(len(adBook.GetEntireAddressBook()))
}

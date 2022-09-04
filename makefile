
prepare:
	go get ./...
	go get github.com/mh-cbon/go-msi
	go install github.com/mh-cbon/go-msi@latest

build-unix:
	
	GOOS=linux 	 GOARCH=amd64 go build -o ./bin/linux/x86/notify-go-lx-x86 ./cmd/main.go 
	GOOS=linux 	 GOARCH=arm64 go build -o ./bin/linux/arm/notify-go-lx-arm ./cmd/main.go
	
	GOOS=darwin  GOARCH=amd64 go build -o ./bin/mac/x86/notify-go-mac-x86 ./cmd/main.go
	GOOS=darwin  GOARCH=arm64 go build -o ./bin/mac/arm/notify-go-mac-arm ./cmd/main.go
	
build-windows:
	go build -o ./bin/windows/x86/notify-go.exe ./cmd/main.go
	go-msi make --msi bin/windows/x86/notify-go.msi --version 0.10 -s wix-assets

prepare:
	go get ./...
	mkdir -p ./bin/cli/linux/x86
	mkdir -p ./bin/cli/linux/arm
	mkdir -p ./bin/cli/mac/x86
	mkdir -p ./bin/cli/mac/arm
	mkdir -p ./bin/cli/windows/x86
	mkdir -p ./bin/cli/windows/arm

build-all:prepare
	
	GOOS=linux 	 GOARCH=amd64 go build -o ./bin/linux/x86/notify-go-lx-x86 ./cmd/main.go 
	GOOS=linux 	 GOARCH=arm64 go build -o ./bin/linux/arm/notify-go-lx-arm ./cmd/main.go
	
	GOOS=darwin  GOARCH=amd64 go build -o ./bin/mac/x86/notify-go-mac-x86 ./cmd/main.go
	GOOS=darwin  GOARCH=arm64 go build -o ./bin/mac/arm/notify-go-mac-arm ./cmd/main.go
	
	GOOS=windows GOARCH=amd64 go build -o ./bin/windows/x86/notify-go-win-x86 ./cmd/main.go
	GOOS=windows GOARCH=arm64 go build -o ./bin/windows/arm/notify-go-win-arm ./cmd/main.go
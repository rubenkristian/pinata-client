build:
	GOOS=windows GOARCH=386 go build -o bin/windows/pinata.exe ./cmd/client/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm/pinata ./cmd/client/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/x64-x86/pinata ./cmd/client/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/x64-x86/pinata ./cmd/client/main.go

run:
	go run ./cmd/client/main.go

example_query:
	go build -o bin/query ./examples/query/main.go

example_pin:
	go build -o bin/pin ./examples/pin/main.go
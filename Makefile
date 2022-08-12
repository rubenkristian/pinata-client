build:
	go build -o bin/client ./cmd/client/main.go

run:
	go run ./cmd/client/main.go

example_query:
	go build -o bin/query ./examples/query/main.go

example_pin:
	go build -o bin/pin ./examples/pin/main.go
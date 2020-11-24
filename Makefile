test:
	go test ./... -v

build:
	go build -o app -ldflags="-X main.Version=$(git rev-parse --short HEAD)" main.go

version := $(shell git rev-parse --short HEAD)
app := actions-test
bin := ${app}-${version}
print-% : ; $(info $($*)) @true

test:
	go test ./... -v

build:
	go build -o ${bin} -ldflags="-X main.Version=${version}" main.go

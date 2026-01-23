.PHONY: build install clean run dev

VERSION := 1.0.0
BINARY := dotfiles-installer
LDFLAGS := -ldflags "-s -w -X main.Version=$(VERSION)"

build:
	go build $(LDFLAGS) -o $(BINARY) ./cmd/installer

install: build
	sudo cp $(BINARY) /usr/local/bin/

clean:
	rm -f $(BINARY)

run:
	go run ./cmd/installer

dev:
	go run ./cmd/installer --dry-run

deps:
	go mod tidy
	go mod download

test:
	go test ./...

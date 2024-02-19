GO_VERSION := 1.18.5
TAG := $(shell git describe --abbrev=0 --tags --always)
HASH := $(shell git rev-parse HEAD)

DATE := $(shell date +%Y-%m-%d.%H:%M:%S)
LDFLAGS := -w -X github.com/ktcf/hello-api/handlers.hash=$(HASH) \
			  -X github.com/ktcf/hello-api/handlers.tag=$(TAG) \
			  -X github.com/ktcf/hello-api/handlers.date=$(DATE)

.PHONY: install-go init-go

setup: install-go init-go install-lint copy-hooks

build:
	go build -ldflags "$(LDFLAGS)" -o api cmd/main.go

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | \
        awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

lint:
	golangci-lint run

install-lint:
	sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.41.1

check-format:
	test -z $$(go fmt ./...)

#TODO add MacOS support
install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf "go$(GO_VERSION).linux-amd64.tar.gz"
	rm "go$(GO_VERSION).linux-amd64.tar.gz"

init-go:
	mkdir -p $(HOME)/go/{bin,src,pkg}
	echo "export GOPATH=$(HOME)/go" >> $(HOME)/.bashrc
	echo "export PATH=\$$PATH:\$$GOPATH/bin:/usr/local/go/bin" >> $(HOME)/.bashrc
	source $(HOME)/.bashrc

copy-hooks:
	chmod +x scripts/hooks/*
	cp -r screipts/hooks/ .git/.
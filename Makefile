GO_VERSION := 1.18

.PHONY: install-go init-go

setup: install-go init-go install-lint

build:
	go build -o api cmd/main.go

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
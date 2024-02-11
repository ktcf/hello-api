GO_VERSION := 1.18

.PHONY: install-go init-go

setup: install-go init-go

build:
	go build -o api cmd/main.go

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

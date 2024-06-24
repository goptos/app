.DEFAULT_GOAL := serve

.PHONY devcontainer: git pre goptos

.PHONY git:
	git config --global core.filemode false && \
	git config --global core.autocrlf true && \
	git config --global init.defaultBranch main

.PHONY pre:
	sudo wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz -P ../ && \
	sudo rm -rf /usr/local/go && \
	sudo tar -C /usr/local -xzf ../go1.22.4.linux-amd64.tar.gz && \
	echo export PATH=\$$PATH:/usr/local/go/bin | tee --append $$HOME/.profile && \
	echo export PATH=\$$PATH:\$$HOME/go/bin | tee --append $$HOME/.profile && \
	. $$HOME/.profile && \
	go version

.PHONY goptos:
	echo export GOPRIVATE=\$$GOPRIVATE,github.com/goptos/runtime | tee --append $$HOME/.profile && \
	echo export GOPRIVATE=\$$GOPRIVATE,github.com/goptos/system | tee --append $$HOME/.profile && \
	echo export GOPRIVATE=\$$GOPRIVATE,github.com/goptos/utils | tee --append $$HOME/.profile && \
	echo export GOPRIVATE=\$$GOPRIVATE,github.com/goptos/lexer | tee --append $$HOME/.profile && \
	echo export GOPRIVATE=\$$GOPRIVATE,github.com/goptos/ast | tee --append $$HOME/.profile && \
	echo export GOPRIVATE=\$$GOPRIVATE,github.com/goptos/view | tee --append $$HOME/.profile && \
	echo export GOPRIVATE=\$$GOPRIVATE,github.com/goptos/cli | tee --append $$HOME/.profile && \
	sudo mkdir -p ../goptos && \
	sudo chmod 777 ../goptos && \
	cd ../goptos && git clone https://github.com/goptos/runtime ; \
	cd ../goptos && git clone https://github.com/goptos/system ; \
	cd ../goptos && git clone https://github.com/goptos/utils ; \
	cd ../goptos && git clone https://github.com/goptos/lexer ; \
	cd ../goptos && git clone https://github.com/goptos/ast ; \
	cd ../goptos && git clone https://github.com/goptos/parser ; \
	cd ../goptos && git clone https://github.com/goptos/cli

.PHONY gogen:
	cd src && \
	go generate

.PHONY build:
	export GOOS=js && \
	export GOARCH=wasm && \
	cd src && \
	go build -o ../dist/main.wasm main.go

.PHONY serve: gogen build
	go run tools/serve/main.go
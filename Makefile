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
	go install github.com/goptos/cli/goptos@latest

.PHONY goptos:
	echo export GONOPROXY=\$$GONOPROXY,github.com/goptos | tee --append $$HOME/.profile && \
	sudo mkdir -p ../goptos && \
	sudo chmod 777 ../goptos && \
	cd ../goptos && git clone https://github.com/goptos/goptos ; \
	cd ../goptos && git clone https://github.com/goptos/stateparser ; \
	cd ../goptos && git clone https://github.com/goptos/runtime ; \
	cd ../goptos && git clone https://github.com/goptos/system ; \
	cd ../goptos && git clone https://github.com/goptos/utils

.PHONY lsp:
	curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash && \
	nvm insall node && \
	nvm use node

.PHONY gogen:
	cd src && \
	go generate

.PHONY gobuild:
	export GOOS=js && \
	export GOARCH=wasm && \
	cd src && \
	go build -o ../dist/main.wasm main.go

.PHONY build:
	goptos build
	goptos package

.PHONY serve:
	goptos serve

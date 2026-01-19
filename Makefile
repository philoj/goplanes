.PHONY: server client

install:
	go install github.com/rakyll/statik@latest

server:
	go run ./server/cmd

statik:
	statik -src=client/assets -dest=client/internal -include=*.jpg,*.png

wasm:
	env GOOS=js GOARCH=wasm go build -o ./client/build/dist/goplanes.wasm github.com/philoj/goplanes/client/cmd
	cp $(shell go env GOROOT)/lib/wasm/wasm_exec.js ./client/build/dist/
	cp ./client/public/index.html ./client/build/dist/

client-image:
	docker build -f ./client/build/Dockerfile ./client/build -t goplanes/frontend:latest
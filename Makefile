.PHONY: server client

server:
	go run ./server/cmd

client:
	go run ./client/cmd 123
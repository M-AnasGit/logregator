package main

import (
	"log"

	"github.com/zhouhui8915/go-socket.io-client"
)

func connectToSocket(key *string) {
	opts := &socketio_client.Options{
		Transport: "socket.io",
		Query:     make(map[string]string),
	}

	keyString := *key

	opts.Query["token"] = keyString
	opts.Query["user"] = "server"
	uri := "https://logregator-server.onrender.com"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("message", func(msg string) {
		log.Printf("on message:%v\n", msg)
	})
	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})
}

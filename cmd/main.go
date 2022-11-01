package main

import (
	f "github.com/ambelovsky/gosf"
	"log"
	k "test/kafka"
)

func init() {

	// Listen on an endpoint
	f.OnConnect(func(client *f.Client, request *f.Request) {
		log.Println("Client connected.")
		client.Join("room")
	})

	f.Listen("new-message", func(client *f.Client, request *f.Request) *f.Message {
		log.Println("New message received.")
		k.Write(request.Message.Text)
		client.Broadcast("room", request.Endpoint, f.NewSuccessMessage(request.Message.Text))
		return f.NewSuccessMessage()
	})

}

func main() {
	go k.Read(func(s string) {
		f.Broadcast("room", "new-message", f.NewSuccessMessage(s))
	})

	f.Startup(map[string]interface{}{
		"port": 9999})
}

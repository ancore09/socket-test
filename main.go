package main

import (
	f "github.com/ambelovsky/gosf"
	"log"
)

func init() {
	// Listen on an endpoint
	f.OnConnect(func(client *f.Client, request *f.Request) {
		log.Println("Client connected.")
		client.Join("room")
	})

	f.Listen("new-message", func(client *f.Client, request *f.Request) *f.Message {
		log.Println("New message received.")
		client.Broadcast("room", request.Endpoint, f.NewSuccessMessage(request.Message.Text))
		return f.NewSuccessMessage()
	})
}

func main() {
	// Start the server using a basic configuration
	f.Startup(map[string]interface{}{
		"port": 9999})
}

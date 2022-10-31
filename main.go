package main

import (
	f "github.com/ambelovsky/gosf"
	"log"
)

func init() {
	// Listen on an endpoint
	f.OnConnect(func(client *f.Client, request *f.Request) {
		log.Println("Client connected.")
	})

	f.Listen("echo", func(client *f.Client, request *f.Request) *f.Message {
		log.Println(request.Message.Text)
		return f.NewSuccessMessage(request.Message.Text)
	})
}

func main() {
	// Start the server using a basic configuration
	f.Startup(map[string]interface{}{
		"port": 9999})
}

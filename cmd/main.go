package main

import (
	f "github.com/ambelovsky/gosf"
	"log"
	"os"
	"strconv"
	k "test/kafka"
)

var (
	m map[string]string
)

func init() {
	m = make(map[string]string)

	// Listen on an endpoint
	f.OnConnect(func(client *f.Client, request *f.Request) {
	})

	f.Listen("register", func(client *f.Client, request *f.Request) *f.Message {
		log.Println(request.Message.GUID)
		m[request.Message.Text] = request.Message.GUID
		return f.NewSuccessMessage("registered")
	})

	f.Listen("join", func(client *f.Client, request *f.Request) *f.Message {
		client.Join(request.Message.Text)
		return f.NewSuccessMessage("joined")
	})

	f.Listen("new-message", func(client *f.Client, request *f.Request) *f.Message {
		log.Println("New message received.")
		k.Write(k.NewMsg(request.Message.Text, request.Message.Body["target"].(string), request.Message.GUID))
		//f.Broadcast(request.Message.GUID, "new-message", f.NewSuccessMessage(request.Message.Text))
		//client.Broadcast("room", request.Endpoint, f.NewSuccessMessage(request.Message.Text))
		return f.NewSuccessMessage()
	})

}

func main() {
	go k.Read("localhost:9092", "messages", func(s k.Msg) {
		msg := f.NewSuccessMessage(s.Text)
		msg.Text = s.Text
		msg.Body = make(map[string]interface{})
		msg.Body["from"] = s.From
		f.Broadcast(s.Target, "new-message", msg)
	})

	p, _ := strconv.Atoi(os.Args[1])
	f.Startup(map[string]interface{}{
		"port": p})
}

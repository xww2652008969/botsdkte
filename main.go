package main

import (
	"botsdkte/plu/pix"
	"github.com/xww2652008969/wbot/MessageType"
	"github.com/xww2652008969/wbot/client"
)

func main() {
	config := client.Clientconfig{
		Wsurl:      "ws://127.0.0.1",
		Wspost:     "3001",
		Wsheader:   nil,
		Clienthttp: "http://127.0.0.1:4000",
	}
	c, err := client.Create(config)
	if err != nil {
		panic(err)
	}
	//c.Register(MessageType.Group, Crazy.Crazy())
	c.Register(MessageType.Group, pix.Getpiximg())
	c.Run()
}

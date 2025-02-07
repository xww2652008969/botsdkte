package main

import (
	botconfig "botsdkte/config"
	"botsdkte/plu/Like"
	"botsdkte/plu/pet"
	"botsdkte/plu/todaywaifu"
	"fmt"
	"github.com/xww2652008969/wbot/client"
)

func main() {
	fmt.Print("正式环境")
	err := botconfig.Read()
	if err != nil {
		panic(err)
	}
	config := client.Clientconfig{
		Wsurl:      botconfig.Gloconfig.Wsurl,
		Wspost:     botconfig.Gloconfig.Wspost,
		Wsheader:   nil,
		Clienthttp: botconfig.Gloconfig.Clienthttp,
	}
	fmt.Println(config)
	c, err := client.Create(config)
	if err != nil {
		panic(err)
	}

	c.RegisterGroupHandle(pet.Pethand())
	c.RegisterGroupHandle(pet.Getcom())
	//c.RegisterGroupHandle(pix.Getpiximg()) 还没设置代理
	c.RegisterGroupHandle(todaywaifu.TodayWaifu())
	c.RegisterGroupHandle(Like.Sedlike())

	c.Run()
}

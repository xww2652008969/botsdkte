package main

import (
	botconfig "botsdkte/config"
	"botsdkte/plu/ff14logs"
	"fmt"
	"github.com/xww2652008969/wbot/client"
)

var idchan chan int64

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
	c, err := client.New(config)
	if err != nil {
		panic(err)
	}
	c.RegisterGroupHandle(ff14logs.FF14log())
	c.Run()
	////c.RegisterGroupHandle(pet.Pethand())
	////c.RegisterGroupHandle(pet.Getcom())
	////c.RegisterGroupHandle(todaywaifu.TodayWaifu())
	////c.RegisterGroupHandle(love.Saylove())
	////c.RegisterGroupHandle(lalfei.Lalafei())
	//c.RegisterGroupHandle(ff14logs.FF14log())
}

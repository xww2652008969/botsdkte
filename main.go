package main

import (
	botconfig "botsdkte/config"
	"botsdkte/plu/hmmt"
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
	c, err := client.New(config)
	if err != nil {
		panic(err)
	}

	//c.RegisterGroupHandle(pet.Pethand())
	//c.RegisterGroupHandle(pet.Getcom())
	//c.RegisterGroupHandle(todaywaifu.TodayWaifu())
	//c.RegisterGroupHandle(love.Saylove())
	//c.RegisterGroupHandle(lalfei.Lalafei())
	c.RegisterGroupHandle(hmmt.SeSe())
	c.Run()
}

//func Athmmt() client.Push {
//	return func(client client.Client) {
//		c := cron.New()
//		_, err := c.AddFunc("*/10 * * * *", func() {
//			client.Sendat(273421673)
//			client.AddText("宝宝来和写挂QAQ")
//			client.SendGroupMsg(853963912)
//		})
//		if err != nil {
//			fmt.Println(err)
//		}
//		c.Start()
//		select {}
//	}
//}

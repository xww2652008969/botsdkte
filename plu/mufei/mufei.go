package mufei

import (
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"log"
	"regexp"
	"time"
)

var ch chan string

func init() {
	ch = make(chan string, 10)
}
func MufeiG() client.Event {
	return func(c client.Client, message client.Message) {
		if message.RawMessage == "母肥" {
			c.AddText("/image 母肥")
			c.SendPrivateMsg(3889045449)
			go func(m client.Message) {
				timeout := time.After(10 * time.Second)
				select {
				case <-timeout:
					fmt.Println("超时")
				case s := <-ch:
					c.Addreply(m.MessageId)
					c.AddImage(s)
					c.SendGroupMsg(m.GroupId)
				}
			}(message)
		}
	}
}
func Mufeip() client.Event {
	return func(c client.Client, message client.Message) {
		if message.UserId != 3889045449 {
			return
		}
		urlPattern := `https?://[^\s)]+`
		re := regexp.MustCompile(urlPattern)
		match := re.FindString(message.RawMessage)
		if match != "" {
			ch <- match
		} else {
			log.Println("没有找到")
		}
	}
}

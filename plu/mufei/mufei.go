package mufei

import (
	"fmt"
	"github.com/xww2652008969/wbot/client/Message"
	"log"
	"regexp"
	"time"
)

var wait bool
var url string

func init() {

}
func MufeiG() Message.Event {
	return func(message Message.Message) {
		if wait {
			return
		}
		if message.RawMessage == "母肥" {
			message.AddText("/image 母肥")
			message.SendPrivateMsg(3889045449)
			wait = true
			for {
				time.Sleep(1 * time.Second)
				if url != "" {
					message.AddImage(url)
					message.SendGroupMsg(message.GroupId)
					url = ""
					break
				}
			}
		}
		wait = false
	}
}
func Mufeip() Message.Event {
	return func(message Message.Message) {
		if message.UserId != 3889045449 && wait == false {
			return
		}
		urlPattern := `https?://[^\s)]+`
		re := regexp.MustCompile(urlPattern)
		match := re.FindString(message.RawMessage)
		if match != "" {
			fmt.Println(match)
			url = match
			wait = false
		} else {
			log.Println("没有找到")
		}
	}
}

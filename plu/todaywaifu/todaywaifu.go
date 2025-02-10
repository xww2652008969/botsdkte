package todaywaifu

import (
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var groupdata waifu
var path string

func init() {
	groupdata = waifu{
		Groupmap: map[int64]group{},
	}
	path, _ = os.Getwd()
	path = filepath.Join(path, "waidata")
	os.MkdirAll(path, 0775)
	groupdata.read()
}
func TodayWaifu() client.Event {
	return func(c client.Client, message client.Message) {
		groupdata.mu.Lock()
		defer groupdata.mu.Unlock()
		groupdata.initdata(c, message).save() //有群聊就初始化
		groupdata.updatedata(c, message).save()
		if message.RawMessage == "今日老婆" {
			dat, err := groupdata.addwaf(message)
			if err != nil {
				c.AddText(err.Error()).SendGroupMsg(message.GroupId)
				groupdata.save()
				return
			}
			c.Addreply(message.MessageId)
			c.AddImage(fmt.Sprintf("https://q.qlogo.cn/g?b=qq&nk=%d&s=640", dat.UserId))
			if dat.Card != "" {
				c.AddText("你的今日老婆是  " + dat.Card)
			} else {
				c.AddText("你的今日老婆是  " + dat.Nickname)
			}
			c.SendGroupMsg(message.GroupId)
			groupdata.save()
		}
		if message.RawMessage == "离婚" {
			c.Addreply(message.MessageId)
			_, err := groupdata.dellwaf(message)
			if err != nil {
				c.AddText(err.Error())
				c.SendGroupMsg(message.GroupId)
				return
			}
			c.AddText("呸渣男")
			c.SendGroupMsg(message.GroupId)
			return
		}
		if message.RawMessage == "换老婆" {
			c.Addreply(message.MessageId)
			_, err := groupdata.dellwaf(message)
			if err != nil {
				c.AddText(err.Error())
				c.SendGroupMsg(message.GroupId)
				return
			}
			dat, err := groupdata.addwaf(message)
			if err != nil {
				c.AddText(err.Error() + "换不了老婆了")
			}
			c.AddImage(fmt.Sprintf("http://q.qlogo.cn/headimg_dl?dst_uin=%d&spec=640&img_type=jpg", dat.UserId))
			if dat.Card != "" {
				c.AddText("你的今日老婆是  " + dat.Card)
			} else {
				c.AddText("你的今日老婆是  " + dat.Nickname)
			}
			c.SendGroupMsg(message.GroupId)
			groupdata.save()
			return
		}
		if message.Message[0].Type != "text" {

		}
		message.Message[0].Data.Text = strings.ReplaceAll(message.Message[0].Data.Text, " ", "")
		if message.Message[0].Data.Text == "娶" && len(message.Message) == 2 {
			if message.Message[1].Type != "at" {
				return
			}
			c.Addreply(message.MessageId)
			useid, _ := strconv.ParseInt(message.Message[1].Data.Qq, 10, 64)
			dat, err := groupdata.sudoaddwaf(message, useid)
			if err.Error() == "你特喵有老婆了，你犯法了" {
				c.AddText(err.Error())
				c.SendGroupMsg(message.GroupId)
				return
			}

			fmt.Println(dat, err)
			c.AddImage(fmt.Sprintf("https://q.qlogo.cn/g?b=qq&nk=%d&s=640", dat.UserId))
			if dat.Card != "" {
				c.AddText("你的今日老婆是  " + dat.Card)
			} else {
				c.AddText("你的今日老婆是  " + dat.Nickname)
			}
			if err != nil && err.Error() != "" {
				c.AddText("\n你这个牛头人")
			}
			c.SendGroupMsg(message.GroupId)
			groupdata.save()
		}
		groupdata.save()
	}
}

package todaywaifu

import (
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"os"
	"path/filepath"
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
			c.AddImage(fmt.Sprintf("http://q.qlogo.cn/headimg_dl?dst_uin=%d&spec=640&img_type=jpg", dat.UserId))
			if dat.Card != "" {
				c.AddText("你的今日老婆是  " + dat.Card)
			} else {
				c.AddText("你的今日老婆是  " + dat.Nickname)
			}
			c.SendGroupMsg(message.GroupId)
			groupdata.save()
		}
	}
}

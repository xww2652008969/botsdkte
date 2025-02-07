package Like

import "github.com/xww2652008969/wbot/client"

func Sedlike() client.Event {
	return func(c client.Client, message client.Message) {
		if (message.RawMessage) == "我要赞大家" && message.UserId == 1271701079 {
			a, _ := c.GetGroupMemberList(message.GroupId)
			for _, v := range a.Data {
				c.SendLike(v.UserId, 10)
			}
		}
	}
}

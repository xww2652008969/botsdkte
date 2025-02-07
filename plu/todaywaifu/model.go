package todaywaifu

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/utils"
	"math/rand"
	"path/filepath"
	"time"
)

type waifu struct {
	Groupmap map[int64]group `json:"groupmap"`
}
type group struct {
	T             time.Time                            `json:"t"`
	Groupuser     map[int64]client.GroupMemberListData `json:"groupuser"`
	Groupuserlist []client.GroupMemberListData         `json:"groupuserlist"`
}

func (w *waifu) initdata(c client.Client, message client.Message) *waifu {
	if _, ok := w.Groupmap[message.GroupId]; ok {
		return w
	} else {
		list, err := c.GetGroupMemberList(message.GroupId)
		if err != nil {
			fmt.Println(err)
			return w
		}
		var g group
		g.Groupuser = make(map[int64]client.GroupMemberListData)
		g.Groupuserlist = list.Data
		g.T = time.Now()
		w.Groupmap[message.GroupId] = g
		fmt.Println(len(g.Groupuserlist))
	}
	return w
}

func (w *waifu) updatedata(c client.Client, message client.Message) *waifu {
	if k, ok := w.Groupmap[message.GroupId]; ok {
		nowtime := time.Now() //现在
		if !isSameDay(nowtime, k.T) {
			list, err := c.GetGroupMemberList(message.GroupId)
			if err != nil {
				fmt.Println(err)
				return w
			}
			var g group
			g.Groupuser = make(map[int64]client.GroupMemberListData)
			g.Groupuserlist = list.Data
			g.T = time.Now()
			w.Groupmap[message.GroupId] = g
		}
	}
	return w
}
func (w *waifu) addwaf(message client.Message) (client.GroupMemberListData, error) {
	if _, ok := w.Groupmap[message.GroupId].Groupuser[message.UserId]; !ok {
		if len(w.Groupmap[message.GroupId].Groupuserlist) == 0 {
			fmt.Println(w.Groupmap[message.GroupId])
			return client.GroupMemberListData{}, errors.New("没有老婆了") //没老婆了
		}
		i := randint(0, len(w.Groupmap[message.GroupId].Groupuserlist)-1)
		w.Groupmap[message.GroupId].Groupuser[message.UserId] = w.Groupmap[message.GroupId].Groupuserlist[i]
		w.dellist(message.GroupId, i)
		return w.Groupmap[message.GroupId].Groupuser[message.UserId], nil
	}
	return w.Groupmap[message.GroupId].Groupuser[message.UserId], nil
}
func (w *waifu) dellist(GroupId int64, i int) {
	g := w.Groupmap[GroupId]
	g.Groupuserlist = append(g.Groupuserlist[:i], g.Groupuserlist[i+1:]...)
	w.Groupmap[GroupId] = g
}
func (w *waifu) read() *waifu {
	da := utils.Readfile(filepath.Join(path, "data.json"))
	if len(da) == 0 {
		da, _ = json.Marshal(w)
		utils.Writefile(filepath.Join(path, "data.json"), da)
	}
	return w
}
func (w *waifu) save() *waifu { //保存数据
	da, err := json.Marshal(w)
	fmt.Println(err)
	utils.Writefile(filepath.Join(path, "data.json"), da)
	return w
}
func randint(min int, max int) int {
	if max == 0 {
		return 1
	}
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}
func isSameDay(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}

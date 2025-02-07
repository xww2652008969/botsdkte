package pet

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	qqbot "github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/utils"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var baseurl string
var client *http.Client
var keyword map[string][]string
var path string

func init() {
	baseurl = "http://192.168.10.209:2333"
	client = &http.Client{}
	keyword = make(map[string][]string)
	getkeylist()
	path, _ = os.Getwd()
	path = path + "/pet/"
	os.Mkdir(path, 0755)
}
func Pethand() qqbot.Event {
	return func(m qqbot.Client, message qqbot.Message) {
		com := message.Message
		if com[0].Type == "reply" {
			com[1].Data.Text = strings.ReplaceAll(com[1].Data.Text, "\n", "")
			com[1].Data.Text = strings.ReplaceAll(com[1].Data.Text, " ", "")
			if _, ok := keyword[com[1].Data.Text]; ok {
				var urlpath []string
				id, _ := strconv.Atoi(com[0].Data.Id)
				msginfo, _ := m.GetMsginfo(id)
				for _, v := range keyword[com[1].Data.Text] {
					pa, err := buildre(v, message, msginfo)
					if err != nil {
						fmt.Println(err)
						continue
					}
					urlpath = append(urlpath, pa)
				}
				if len(urlpath) == 0 {
					return
				}
				m.Addreply(message.MessageId)
				for _, v := range urlpath {
					m.AddImage(v)
				}
				m.SendGroupMsg(message.GroupId)
				return
			}
		}
		if com[0].Type != "text" {
			return
		}
		com[0].Data.Text = strings.ReplaceAll(com[0].Data.Text, "\n", "")
		if _, ok := keyword[com[0].Data.Text]; ok {
			var urlpath []string
			for _, v := range keyword[com[0].Data.Text] {
				pa, err := buildreq(v, message)
				if err != nil {
					fmt.Println(err)
					continue
				}
				urlpath = append(urlpath, pa)
			}
			if len(urlpath) == 0 {
				return
			}
			m.Addreply(message.MessageId)
			for _, v := range urlpath {
				m.AddImage(v)
			}
			m.SendGroupMsg(message.GroupId)
			return
		}
	}
}
func Getcom() qqbot.Event {
	return func(c qqbot.Client, message qqbot.Message) {
		if message.RawMessage == "指令" {
			var t string
			for k, _ := range keyword {
				t += k + "   "
			}
			c.Addreply(message.MessageId)
			c.AddText(t)
			c.SendGroupMsg(message.GroupId)
		}
	}
}

func buildreq(id string, message qqbot.Message) (string, error) {
	var petreq reqstruct
	petreq.Id = id
	petreq.Image.From = fmt.Sprintf("http://q.qlogo.cn/headimg_dl?dst_uin=%d&spec=640&img_type=jpg", message.UserId)
	for _, v := range message.Message {
		if v.Type == "at" {
			petreq.Image.To = fmt.Sprintf("http://q.qlogo.cn/headimg_dl?dst_uin=%s&spec=640&img_type=jpg", v.Data.Qq)
			break
		}
		if v.Type == "image" {
			petreq.Image.To = fmt.Sprintf(v.Data.Url)
			break
		}
		petreq.Image.To = petreq.Image.From
	}
	petreq.Text.To = "te"
	fmt.Println(petreq)
	jsondata, _ := json.Marshal(petreq)
	hettpreq, _ := http.NewRequest("POST", baseurl+"/generate", bytes.NewBuffer(jsondata))
	res, err := client.Do(hettpreq)
	if err != nil {
		return "", errors.New("有问题")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		data, _ := io.ReadAll(res.Body)
		fmt.Println(string(data))
		return "", errors.New(res.Status)
	}
	data, _ := io.ReadAll(res.Body)
	md5 := utils.Gedmd5(data)
	utils.Writefile(path+md5, data)
	return path + md5, nil
}
func buildre(id string, message qqbot.Message, m qqbot.MsgInfo) (string, error) {
	var petreq reqstruct
	petreq.Id = id
	petreq.Image.From = fmt.Sprintf("http://q.qlogo.cn/headimg_dl?dst_uin=%d&spec=640&img_type=jpg", message.UserId)
	fmt.Println(m)
	for _, v := range m.Data.Message {
		fmt.Println(v)
		if v.Type == "image" {
			fmt.Println(v.Data)
			petreq.Image.To = fmt.Sprintf(v.Data.Url)
			break
		}
		petreq.Image.To = petreq.Image.From
	}
	petreq.Text.To = "te"
	fmt.Println(petreq)
	jsondata, _ := json.Marshal(petreq)
	hettpreq, _ := http.NewRequest("POST", baseurl+"/generate", bytes.NewBuffer(jsondata))
	res, err := client.Do(hettpreq)
	if err != nil {
		return "", errors.New("有问题")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		data, _ := io.ReadAll(res.Body)
		fmt.Println(string(data))
		return "", errors.New(res.Status)
	}
	data, _ := io.ReadAll(res.Body)
	md5 := utils.Gedmd5(data)
	utils.Writefile(path+md5, data)
	return path + md5, nil
}

func getkeylist() error {
	var keylist key
	req, _ := http.NewRequest("GET", baseurl+"/", nil)
	res, _ := client.Do(req)
	if res.StatusCode != 200 {
		return errors.New("获取key失败不启动")
	}
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &keylist)
	for _, v := range keylist.Templates {
		for _, v1 := range v.Metadata.Alias {
			if _, ok := keyword[v1]; ok {
				keyword[v1] = append(keyword[v1], v.Id)
			} else {
				keyword[v1] = make([]string, 0)
				keyword[v1] = []string{v.Id}
			}
		}
	}
	return nil
}

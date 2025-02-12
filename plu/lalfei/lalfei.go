package lalfei

import (
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var path string
var imgpath []string

func init() {
	imgpath = make([]string, 0)
	root := "lalafell" // 替换为你的目录路径
	path, _ = os.Getwd()
	// 使用 filepath.Walk 函数遍历目录
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是文件
		if !info.IsDir() {
			imgpath = append(imgpath, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
func Lalafei() client.Event {
	return func(client client.Client, message client.Message) {
		com := strings.Split(message.RawMessage, " ")
		if len(com) == 1 {
			if com[0] == "母肥" {
				client.Addreply(message.MessageId)
				client.AddText("谢谢HMMT\n")
				i := randint(0, len(imgpath))
				client.AddImage(filepath.Join(path, imgpath[i]))
				client.SendGroupMsg(message.GroupId)
				return
			}
		}
		if len(com) == 2 {
			if com[0] == "母肥" {
				num, err := strconv.Atoi(com[1])
				if err != nil {
					return
				}
				if num > 3 {
					client.Addreply(message.MessageId)
					client.AddText("你好贪心哦")
					client.SendGroupMsg(message.GroupId)
					return
				}
				client.Addreply(message.MessageId)
				client.AddText("谢谢HMMT\n")
				for i := 0; i < num; i++ {
					time.Sleep(1)
					o := randint(0, len(imgpath))
					client.AddImage(filepath.Join(path, imgpath[o]))
				}
				client.SendGroupMsg(message.GroupId)
				return
			}
		}
	}
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

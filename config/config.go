package botconfig

import (
	"encoding/json"
	"errors"
	"github.com/xww2652008969/wbot/client/utils"
)

type config struct {
	Wsurl      string `json:"wsurl"`
	Wspost     string `json:"wspost"`
	Wsheader   string `json:"wsheader"`
	Clienthttp string `json:"clienthttp"`
	Petbaseurl string `json:"petbaseurl"`
}

var Gloconfig config

func Read() error {
	data := utils.Readfile("config.json")
	if len(data) > 0 {
		json.Unmarshal(data, &Gloconfig)
		return nil
	}
	var d config
	da, _ := json.Marshal(d)
	utils.Writefile("config.json", da)
	return errors.New("第一次启动需要修改配置文件")
}

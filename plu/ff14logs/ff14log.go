package ff14logs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/utils"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/freetype"
)

var path string
var zones []Zone
var apiUrl string
var oauthurl string
var ouatdata Ouatdata
var token Token
var font *truetype.Font
var ico Bossicon

func init() {
	path, _ = os.Getwd()
	path = filepath.Join(path, "data", "ff14logs")
	ico.bossicon = make(map[int]image.Image)
	ico.jobicon = make(map[string]image.Image)
	os.Mkdir(path, 0775) //直接创建文件夹
	zones = []Zone{
		{Id: 54, Name: "万魔殿 荒天之狱", Difficulty: 101},
		{Id: 49, Name: "万魔殿 炼净之狱", Difficulty: 101},
		{Id: 44, Name: "万魔殿 边境之狱", Difficulty: 101},
		{Id: 62, Name: "阿卡狄亚竞技场 轻量级", Difficulty: 101},
		{Id: 53, Name: "欧米茄绝境验证战", Difficulty: 100},
		{Id: 45, Name: "幻想龙诗绝境战", Difficulty: 100},
		{Id: 43, Name: "绝境战（旧版本）", Difficulty: 100},
	}
	apiUrl = "https://cn.fflogs.com/api/v2/client"
	oauthurl = "https://cn.fflogs.com/oauth/token"
	ouatdata = Ouatdata{
		ClientId:     "9e37515c-624a-4913-b710-95886ecda1aa",
		ClientSecret: "QRWYJsYCPNEbHgZAUOiV5SKohFeiSyB2VAHHaCDT",
		GrantType:    "client_credentials",
	}
	da, _ := json.Marshal(ouatdata)
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "*/*",
		"User-Agent":   "Apifox/1.0.0 (https://apifox.com)",
	}
	res, _ := utils.Httppost(oauthurl, headers, bytes.NewBuffer(da))
	da, _ = io.ReadAll(res.Body)
	json.Unmarshal(da, &token)
	fontdata := utils.Readfile(filepath.Join(path, "FZMiaoWuK.TTF"))
	font, _ = freetype.ParseFont(fontdata)
}
func FF14log() client.Event {
	return func(client client.Client, message client.Message) {
		com := strings.Split(message.RawMessage, " ")
		if len(com) != 3 {
			return
		}
		if com[0] != "fflogs" {
		}
		character_name := com[2]
		server := com[1]
		fmt.Println(server, character_name)
		client.Addreply(message.MessageId)
		var raninfodata []RankingInfo
		for _, v := range zones {
			data, err := getdata(server, character_name, v.Id, v.Difficulty)
			if err != nil {

				continue
			}
			parseResponseData(data, &raninfodata)
		}
		if len(raninfodata) == 0 {
			return
		}
		s := generateImage(character_name, server, raninfodata)
		client.AddImage(s)
		client.SendGroupMsg(message.GroupId)
	}
}

func getdata(server, name string, zoneID, difficulty int) (FF14Logdata, error) {
	query := fmt.Sprintf(`
		query {
			characterData {
				character(name: "%s", serverRegion: "cn", serverSlug: "%s") {
					zoneRankings(zoneID: %d, difficulty: %d)
				}
			}
		}
	`, name, server, zoneID, difficulty)
	headers := map[string]string{
		"Authorization": "Bearer " + token.AccessToken,
		"Content-Type":  "application/json",
		"Accept":        "*/*",
		"User-Agent":    "Apifox/1.0.0 (https://apifox.com)",
	}
	da := GraphqlQuery{Query: query}
	postdada, _ := json.Marshal(da)
	res, _ := utils.Httppost(apiUrl, headers, bytes.NewBuffer(postdada))
	s, _ := io.ReadAll(res.Body)
	var j FF14Logdata
	if len(s) < 100 {
		return j, errors.New("未找到玩家数据")
	}
	json.Unmarshal(s, &j)
	return j, nil
}
func parseResponseData(data FF14Logdata, outdata *[]RankingInfo) {
	rankings := data.Data.CharacterData.Character.ZoneRankings.Rankings
	for _, v := range rankings {
		if v.Spec == "" {
			continue
		}
		encounter_id := v.Encounter.Id
		encounter_name := v.Encounter.Name
		rank_percent := v.RankPercent
		total_kills := v.TotalKills
		spec := v.Spec
		best_amount := v.BestAmount
		var a = RankingInfo{
			encounterId:   encounter_id,
			encounterName: encounter_name,
			rankPercent:   rank_percent,
			totalKills:    total_kills,
			spec:          spec,
			bestAmount:    best_amount,
		}
		*outdata = append(*outdata, a)
	}
}
func generateImage(name, server string, data []RankingInfo) string {
	width := 640
	heigjt := 170 + len(data)*80

	outimg := image.NewRGBA(image.Rect(0, 0, width, heigjt))
	draw.Draw(outimg, outimg.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src) //覆盖背景
	//绘制标题
	drawtext(outimg, getpos(10, 60), fmt.Sprintf("%s-%s", name, server), color.White)
	drawtext(outimg, getpos(20, 130), "Boss", getcoler(180, 189, 255))
	drawtext(outimg, getpos(290, 130), "Best", getcoler(180, 189, 255))
	drawtext(outimg, getpos(410, 130), "Highest RDPS", getcoler(180, 189, 255))
	drawtext(outimg, getpos(590, 130), "Kills", getcoler(180, 189, 255))

	y_offset := 160 //偏移
	for _, v := range data {
		drawrankinfo(outimg, v, y_offset)
		y_offset += 80
	}
	f, _ := os.Create(filepath.Join(path, name+server))
	png.Encode(f, outimg)
	f.Close()
	return filepath.Join(path, name+server)
}
func drawtext(img *image.RGBA, pos image.Point, text string, textcolor color.Color) error {
	fontSize := float64(24)
	dpi := float64(72)
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(textcolor))
	opts := truetype.Options{Size: fontSize, DPI: dpi}
	face := truetype.NewFace(font, &opts)
	defer face.Close()
	metrics := face.Metrics()
	ascend := metrics.Ascent.Ceil()
	pt := freetype.Pt(pos.X, pos.Y+ascend)
	_, err := c.DrawString(text, pt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func drawrankinfo(outimg *image.RGBA, data RankingInfo, off int) {
	//绘制boss 图标
	bosico := ico.geticon(data.encounterId)
	if bosico == nil {
		fmt.Println("获取错误")
		return
	}
	drawRectangle := image.Rect(10, off+8, 10+bosico.Bounds().Dx(), off+8+bosico.Bounds().Dy())
	draw.Draw(outimg, drawRectangle, bosico, getpos(0, 0), draw.Over)
	drawtext(outimg, getpos(84, off+20), data.encounterName, getcoler(180, 189, 255))

	//绘制rankPercent
	rank_color := getRankColor(data.rankPercent)
	drawtext(outimg, getpos(290, off+20), string(strconv.Itoa(int(data.rankPercent))), rank_color)

	//绘制职业图标
	jobico := ico.getjobicon(data.spec)
	if jobico == nil {
		fmt.Println("获取职业图标错误")
		return
	}
	drawRectangle = image.Rect(320, off+20, 320+jobico.Bounds().Dx(), off+20+jobico.Bounds().Dy())
	draw.Draw(outimg, drawRectangle, jobico, getpos(0, 0), draw.Over)
	//# 绘制 totalKills 和 bestAmount
	drawtext(outimg, getpos(610, off+20), strconv.Itoa(data.totalKills), getcoler(225, 242, 245))
	drawtext(outimg, getpos(450, off+20), fmt.Sprintf("%.2f", data.bestAmount), getcoler(180, 189, 255))
}
func getpos(x, y int) image.Point {
	return image.Point{
		X: x,
		Y: y,
	}
}
func getcoler(r, g, b uint8) color.RGBA {
	a := color.RGBA{R: r, G: g, B: b, A: 0}
	return a
}

func getRankColor(rank float64) color.RGBA {
	if rank == -1 {
		return getcoler(128, 128, 128) // 灰色
	}
	if rank == 100 {
		return getcoler(229, 204, 128) // 柔和的黄色
	}
	if rank >= 99 {
		return getcoler(226, 104, 168) // 粉色
	}
	if rank >= 95 {
		return getcoler(255, 128, 0) // 橙色
	}
	if rank >= 75 {
		return getcoler(163, 53, 238) // 紫色
	}
	if rank >= 50 {
		return getcoler(0, 112, 255) // 蓝色
	}
	if rank >= 25 {
		return getcoler(30, 255, 0) // 绿色
	}
	return getcoler(128, 128, 128) // 默认灰色
}

//func Te() {
//	j, err := getdata("延夏", "露露米", 62, 101)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(parseResponseData(j))
//}

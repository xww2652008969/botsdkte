package feitian

import (
	"bytes"
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/utils"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var path string
var rng *rand.Rand

func init() {
	path, _ = os.Getwd()
	path = filepath.Join(path, "data", "feitian")
	os.MkdirAll(path, 0755)
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}
func Feitian() client.Event {
	return func(client client.Client, message client.Message) {
		if len(message.Message) == 2 {
			com := strings.ReplaceAll(message.Message[0].Data.Text, " ", "")
			if com == "飞天" && message.Message[1].Type == "at" {
				bg, err := getqqimg(message.Message[1].Data.Qq)
				if err != nil {
					return
				}
				d := utils.Readfile(filepath.Join(path, "data.png"))
				f, _ := png.Decode(bytes.NewReader(d))
				var wg sync.WaitGroup
				for k := 1; k <= 50; k++ {
					wg.Add(1)
					go func(b *image.RGBA) {
						defer wg.Done()
						tianbeijing(b, rotateImage(f))
					}(bg)
				}
				wg.Wait()
				out := filepath.Join(path, fmt.Sprintf("%s.jpg", message.Message[1].Data.Qq))
				opt, _ := os.Create(out)
				jpeg.Encode(opt, bg, nil)
				opt.Close()
				client.Addreply(message.MessageId)
				client.AddImage(out)
				client.SendGroupMsg(message.GroupId)
			}
		}
	}
}

func rotateImage(img image.Image) *image.RGBA {
	// 计算旋转的弧度
	angle := rng.Float64() * 360
	radians := angle * math.Pi / 180

	// 获取图像的尺寸
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// 创建一个新的 RGBA 图像
	rotated := image.NewRGBA(bounds)

	// 计算旋转的中心
	centerX, centerY := float64(width)/2, float64(height)/2

	// 遍历每个像素并计算新位置
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			// 计算相对于中心的坐标
			oldX := float64(x) - centerX
			oldY := float64(y) - centerY

			// 应用旋转矩阵
			newX := int(oldX*math.Cos(radians) - oldY*math.Sin(radians) + centerX)
			newY := int(oldX*math.Sin(radians) + oldY*math.Cos(radians) + centerY)

			// 检查新坐标是否在范围内
			if newX >= 0 && newX < width && newY >= 0 && newY < height {
				rotated.Set(newX, newY, img.At(x, y))
			}
		}
	}

	return rotated
}
func tianbeijing(outimg *image.RGBA, img image.Image) {
	x := rng.Intn(640)
	y := rng.Intn(640)
	position := image.Point{X: x, Y: y}
	a := rotateImage(img)
	drawRectangle := image.Rect(position.X, position.Y, position.X+a.Bounds().Dx(), position.Y+a.Bounds().Dy())
	draw.Draw(outimg, drawRectangle, a, image.Point{0, 0}, draw.Over)
}
func getqqimg(qq string) (*image.RGBA, error) {
	url := fmt.Sprintf("https://q.qlogo.cn/g?b=qq&nk=%s&s=640", qq)
	res, err := utils.Httpget(url, nil)
	if err != nil {
		fmt.Println(err)
		return &image.RGBA{}, err
	}
	bg, err := jpeg.Decode(res.Body)
	if err != nil {
		fmt.Println(err)
		return &image.RGBA{}, err
	}
	outimg := image.NewRGBA(bg.Bounds())
	draw.Draw(outimg, bg.Bounds(), bg, image.Point{0, 0}, draw.Src)
	return outimg, nil
}

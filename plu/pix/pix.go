package pix

import (
	"encoding/json"
	"fmt"
	"github.com/xww2652008969/wbot/client"
	"github.com/xww2652008969/wbot/client/utils"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var lastPage map[string]int
var exclude []string

type Search struct {
	Error bool `json:"error"`
	Body  struct {
		IllustManga struct {
			Data []struct {
				Id                      string      `json:"id"`
				Title                   string      `json:"title"`
				IllustType              int         `json:"illustType"`
				XRestrict               int         `json:"xRestrict"`
				Restrict                int         `json:"restrict"`
				Sl                      int         `json:"sl"`
				Url                     string      `json:"url"`
				Description             string      `json:"description"`
				Tags                    []string    `json:"tags"`
				UserId                  string      `json:"userId"`
				UserName                string      `json:"userName"`
				Width                   int         `json:"width"`
				Height                  int         `json:"height"`
				PageCount               int         `json:"pageCount"`
				IsBookmarkable          bool        `json:"isBookmarkable"`
				BookmarkData            interface{} `json:"bookmarkData"`
				Alt                     string      `json:"alt"`
				TitleCaptionTranslation struct {
					WorkTitle   interface{} `json:"workTitle"`
					WorkCaption interface{} `json:"workCaption"`
				} `json:"titleCaptionTranslation"`
				CreateDate      time.Time `json:"createDate"`
				UpdateDate      time.Time `json:"updateDate"`
				IsUnlisted      bool      `json:"isUnlisted"`
				IsMasked        bool      `json:"isMasked"`
				AiType          int       `json:"aiType"`
				ProfileImageUrl string    `json:"profileImageUrl"`
			} `json:"data"`
			Total          int `json:"total"`
			LastPage       int `json:"lastPage"`
			BookmarkRanges []struct {
				Min *int        `json:"min"`
				Max interface{} `json:"max"`
			} `json:"bookmarkRanges"`
		} `json:"illustManga"`
		Popular struct {
			Recent []struct {
				Id                      string      `json:"id"`
				Title                   string      `json:"title"`
				IllustType              int         `json:"illustType"`
				XRestrict               int         `json:"xRestrict"`
				Restrict                int         `json:"restrict"`
				Sl                      int         `json:"sl"`
				Url                     string      `json:"url"`
				Description             string      `json:"description"`
				Tags                    []string    `json:"tags"`
				UserId                  string      `json:"userId"`
				UserName                string      `json:"userName"`
				Width                   int         `json:"width"`
				Height                  int         `json:"height"`
				PageCount               int         `json:"pageCount"`
				IsBookmarkable          bool        `json:"isBookmarkable"`
				BookmarkData            interface{} `json:"bookmarkData"`
				Alt                     string      `json:"alt"`
				TitleCaptionTranslation struct {
					WorkTitle   interface{} `json:"workTitle"`
					WorkCaption interface{} `json:"workCaption"`
				} `json:"titleCaptionTranslation"`
				CreateDate      time.Time `json:"createDate"`
				UpdateDate      time.Time `json:"updateDate"`
				IsUnlisted      bool      `json:"isUnlisted"`
				IsMasked        bool      `json:"isMasked"`
				AiType          int       `json:"aiType"`
				ProfileImageUrl string    `json:"profileImageUrl"`
			} `json:"recent"`
			Permanent []struct {
				Id                      string      `json:"id"`
				Title                   string      `json:"title"`
				IllustType              int         `json:"illustType"`
				XRestrict               int         `json:"xRestrict"`
				Restrict                int         `json:"restrict"`
				Sl                      int         `json:"sl"`
				Url                     string      `json:"url"`
				Description             string      `json:"description"`
				Tags                    []string    `json:"tags"`
				UserId                  string      `json:"userId"`
				UserName                string      `json:"userName"`
				Width                   int         `json:"width"`
				Height                  int         `json:"height"`
				PageCount               int         `json:"pageCount"`
				IsBookmarkable          bool        `json:"isBookmarkable"`
				BookmarkData            interface{} `json:"bookmarkData"`
				Alt                     string      `json:"alt"`
				TitleCaptionTranslation struct {
					WorkTitle   interface{} `json:"workTitle"`
					WorkCaption interface{} `json:"workCaption"`
				} `json:"titleCaptionTranslation"`
				CreateDate      time.Time `json:"createDate"`
				UpdateDate      time.Time `json:"updateDate"`
				IsUnlisted      bool      `json:"isUnlisted"`
				IsMasked        bool      `json:"isMasked"`
				AiType          int       `json:"aiType"`
				ProfileImageUrl string    `json:"profileImageUrl"`
			} `json:"permanent"`
		} `json:"popular"`
		RelatedTags []string `json:"relatedTags"`
		ZoneConfig  struct {
			Header struct {
				Url string `json:"url"`
			} `json:"header"`
			Footer struct {
				Url string `json:"url"`
			} `json:"footer"`
			Infeed struct {
				Url string `json:"url"`
			} `json:"infeed"`
			Logo struct {
				Url string `json:"url"`
			} `json:"logo"`
			AdLogo struct {
				Url string `json:"url"`
			} `json:"ad_logo"`
		} `json:"zoneConfig"`
		ExtraData struct {
			Meta struct {
				Title              string `json:"title"`
				Description        string `json:"description"`
				Canonical          string `json:"canonical"`
				AlternateLanguages struct {
					Ja string `json:"ja"`
					En string `json:"en"`
				} `json:"alternateLanguages"`
				DescriptionHeader string `json:"descriptionHeader"`
			} `json:"meta"`
		} `json:"extraData"`
	} `json:"body"`
}
type Illust struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		IllustId      string    `json:"illustId"`
		IllustTitle   string    `json:"illustTitle"`
		IllustComment string    `json:"illustComment"`
		Id            string    `json:"id"`
		Title         string    `json:"title"`
		Description   string    `json:"description"`
		IllustType    int       `json:"illustType"`
		CreateDate    time.Time `json:"createDate"`
		UploadDate    time.Time `json:"uploadDate"`
		Restrict      int       `json:"restrict"`
		XRestrict     int       `json:"xRestrict"`
		Sl            int       `json:"sl"`
		Urls          struct {
			Mini     string `json:"mini"`
			Thumb    string `json:"thumb"`
			Small    string `json:"small"`
			Regular  string `json:"regular"`
			Original string `json:"original"`
		} `json:"urls"`
		Tags struct {
			AuthorId string `json:"authorId"`
			IsLocked bool   `json:"isLocked"`
			Tags     []struct {
				Tag         string `json:"tag"`
				Locked      bool   `json:"locked"`
				Deletable   bool   `json:"deletable"`
				UserId      string `json:"userId,omitempty"`
				Translation struct {
					En string `json:"en"`
				} `json:"translation,omitempty"`
				UserName string `json:"userName,omitempty"`
			} `json:"tags"`
			Writable bool `json:"writable"`
		} `json:"tags"`
		Alt                  string        `json:"alt"`
		UserId               string        `json:"userId"`
		UserName             string        `json:"userName"`
		UserAccount          string        `json:"userAccount"`
		LikeData             bool          `json:"likeData"`
		Width                int           `json:"width"`
		Height               int           `json:"height"`
		PageCount            int           `json:"pageCount"`
		BookmarkCount        int           `json:"bookmarkCount"`
		LikeCount            int           `json:"likeCount"`
		CommentCount         int           `json:"commentCount"`
		ResponseCount        int           `json:"responseCount"`
		ViewCount            int           `json:"viewCount"`
		BookStyle            interface{}   `json:"bookStyle"`
		IsHowto              bool          `json:"isHowto"`
		IsOriginal           bool          `json:"isOriginal"`
		ImageResponseOutData []interface{} `json:"imageResponseOutData"`
		ImageResponseData    []interface{} `json:"imageResponseData"`
		ImageResponseCount   int           `json:"imageResponseCount"`
		PollData             interface{}   `json:"pollData"`
		SeriesNavData        interface{}   `json:"seriesNavData"`
		DescriptionBoothId   interface{}   `json:"descriptionBoothId"`
		DescriptionYoutubeId interface{}   `json:"descriptionYoutubeId"`
		ComicPromotion       interface{}   `json:"comicPromotion"`
		FanboxPromotion      interface{}   `json:"fanboxPromotion"`
		ContestBanners       []interface{} `json:"contestBanners"`
		IsBookmarkable       bool          `json:"isBookmarkable"`
		BookmarkData         interface{}   `json:"bookmarkData"`
		ContestData          interface{}   `json:"contestData"`
		ZoneConfig           struct {
			Responsive struct {
				Url string `json:"url"`
			} `json:"responsive"`
			Rectangle struct {
				Url string `json:"url"`
			} `json:"rectangle"`
			X500 struct {
				Url string `json:"url"`
			} `json:"500x500"`
			Header struct {
				Url string `json:"url"`
			} `json:"header"`
			Footer struct {
				Url string `json:"url"`
			} `json:"footer"`
			ExpandedFooter struct {
				Url string `json:"url"`
			} `json:"expandedFooter"`
			Logo struct {
				Url string `json:"url"`
			} `json:"logo"`
			AdLogo struct {
				Url string `json:"url"`
			} `json:"ad_logo"`
			Relatedworks struct {
				Url string `json:"url"`
			} `json:"relatedworks"`
		} `json:"zoneConfig"`
		ExtraData struct {
			Meta struct {
				Title              string `json:"title"`
				Description        string `json:"description"`
				Canonical          string `json:"canonical"`
				AlternateLanguages struct {
					Ja string `json:"ja"`
					En string `json:"en"`
				} `json:"alternateLanguages"`
				DescriptionHeader string `json:"descriptionHeader"`
				Ogp               struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Type        string `json:"type"`
				} `json:"ogp"`
				Twitter struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Card        string `json:"card"`
				} `json:"twitter"`
			} `json:"meta"`
		} `json:"extraData"`
		TitleCaptionTranslation struct {
			WorkTitle   interface{} `json:"workTitle"`
			WorkCaption interface{} `json:"workCaption"`
		} `json:"titleCaptionTranslation"`
		IsUnlisted           bool        `json:"isUnlisted"`
		Request              interface{} `json:"request"`
		CommentOff           int         `json:"commentOff"`
		AiType               int         `json:"aiType"`
		ReuploadDate         interface{} `json:"reuploadDate"`
		LocationMask         bool        `json:"locationMask"`
		CommissionLinkHidden bool        `json:"commissionLinkHidden"`
		IsLoginOnly          bool        `json:"isLoginOnly"`
	} `json:"body"`
}

var path string
var dirPath string

func init() {
	lastPage = make(map[string]int)
	exclude = []string{"NovelAI", "ProjectSekai", "fuury", "AIイラスト", "AI生成"}
	path, _ = os.Getwd()
	dirPath = filepath.Join(path, "img")
	err := os.Mkdir(dirPath, 0755) // 0755 是目录的权限
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("Directory already exists:", dirPath)
		} else {
			fmt.Println("Error creating directory:", err)
			return
		}
	} else {
		fmt.Println("Directory created successfully:", dirPath)
	}
}
func Getpiximg() client.Event {
	return func(c client.Client, message client.Message) {
		com := strings.Split(message.RawMessage, " ")
		fmt.Println(len(com))
		if len(com) < 2 {
			return
		}
		if com[0] == "搜索图片" {
			p := 1
			if k, ok := lastPage[com[1]]; ok {
				fmt.Println(lastPage)
				p = Randint(1, k)
			} else {
				p = 1
			}
			a, err := sr18(com[1], false, p)
			if err != nil {
				return
			}
			lastPage[com[1]] = a.Body.IllustManga.LastPage
			fmt.Println(a.Body.IllustManga.LastPage)
			if len(a.Body.IllustManga.Data) < 1 {
				return
			}
			i := Randint(0, len(a.Body.IllustManga.Data))
			list, err := getIllust(a.Body.IllustManga.Data[i].Id)
			if err != nil {
				fmt.Println(err)
				return
			}
			s := getimg(list.Body.Urls.Original)
			fmt.Println(s)
			if len(s) > 5 {
				c.Addreply(message.MessageId)
				c.AddImage(s)
				c.SendGroupMsg(message.GroupId)
			}
			return
		}
	}
}

func Randint(min int, max int) int {
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

// 传入图片地址 获取图片
func getimg(url string) string {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("accept", " text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("priority", "u=0, i")
	req.Header.Add("referer", "https://www.pixiv.net/")
	req.Header.Add("sec-ch-ua", `"Not A(Brand";v="8", "Chromium";v="132", "Microsoft Edge";v="132"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "cross-site")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", " 1")
	req.Header.Add("user-agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")
	req.Header.Add("Host", "www.pixiv.net")
	fmt.Println(req.Header)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if len(body) > 1000 {
		s := utils.Gedmd5(body)
		w := filepath.Join(dirPath, s)
		utils.Writefile(w, body)
		return w
	}
	return ""
}
func sr18(key string, r bool, p int) (Search, error) {
	// 创建一个新的请求
	var s Search
	//s_tag 为全年龄  r18为r18
	baseURL := "https://www.pixiv.net/ajax/search/artworks/" + url.QueryEscape(key)
	queryParams := url.Values{}
	queryParams.Set("word", key)
	queryParams.Set("order", "date_d")
	if r {
		queryParams.Set("mode", "r18")
	} else {
		queryParams.Set("mode", "safe")
	}
	queryParams.Set("p", strconv.Itoa(p))
	queryParams.Set("csw", "0")
	queryParams.Set("s_mode", "s_tag")
	queryParams.Set("type", "all")
	queryParams.Set("lang", "zh")
	queryParams.Set("version", "1514cd4826094c32a804b4de6def5f2209963922")
	fullURL := baseURL + "?" + queryParams.Encode()
	fmt.Println(fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return s, err
	}

	// 设置请求头
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Baggage", "sentry-environment=production,sentry-release=1514cd4826094c32a804b4de6def5f2209963922,sentry-public_key=7b15ebdd9cf64efb88cfab93783df02a,sentry-trace_id=96c04fd1748e43b5b58a8cc1007eef1c,sentry-sample_rate=0.0001")
	req.Header.Set("Cache-Control", "max-age=0")
	//先不设置cookie
	req.Header.Set("Cookie", "first_visit_datetime_pc=2024-10-27%2021%3A31%3A40; p_ab_id=2; p_ab_id_2=4; p_ab_d_id=692958056; yuid_b=QHl0dBI; c_type=25; privacy_policy_notification=0; a_type=0; b_type=1; PHPSESSID=40581705_rQmgRA1guZzvkWGVwysacRwWo2CcGh2C; device_token=6a83fbbbdb3dc55a762744511e49b9ca; _ga_MZ1NL4PHH0=GS1.1.1736429551.2.0.1736429573.0.0.0; privacy_policy_agreement=0; login_ever=yes; _gcl_au=1.1.987029509.1737216645; _gid=GA1.2.1484901853.1738498405; __cf_bm=a4AX2lqOqLQQbrzPiG4d0a_LbWYxLR6tqY3hYdQrjS4-1738500229-1.0.1.1-2VYY.ZeKUu1XGL78D2HRj1mLRrr7lakW1Y8kOcKFSvRS.cJr1_I5jpbMQ5nRH.UAtiSwx2vAwMIhwe1z8oijmRZimOamtsn7mpJkua8bN9g; cf_clearance=gY8jZtdSlKveose00BmumlYWbgDaY1ZxPVoYfO33.NE-1738500230-1.2.1.1-SVz6XzrRYK4ajt7dL2iwXuRXtxFJIF8_6zFF35uXzLtL1PCdrboRG7qfim_b3j8kHxnWkmbsVqyOMjNuzJ5e0RZgkGJLJ2yl9kXGVs44.x.zE8KwUC.3sRduMnhAaDaE9_6SJyfKQXgRy.Y.T2KSPharG3EitMyExnOlNKvspL3EDXHLYktRN4.s0227UFDm73et4XVqBj1l3jOQ.ZVbHyBBIBiMwUDW3yZVz7PBS5_ry3TFL1ajDyHgm7mPnWLOxnt3RxHZ4eLaxPtCRfnkQGrLb0I1Ep7kG5cVaXMlPEw; _ga=GA1.2.1094039263.1730032302; _ga_75BBYNYN9J=GS1.1.1738498401.3.1.1738500234.0.0.0")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://www.pixiv.net/tags/"+url.QueryEscape(key)+"/artworks?mode=r18&s_mode=s_tag") //和搜索地址有关
	req.Header.Set("Sec-CH-UA", `"Not A(Brand";v="8", "Chromium";v="132", "Microsoft Edge";v="132"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	//req.Header.Set("Sentry-Trace", "96c04fd1748e43b5b58a8cc1007eef1c-b0200e41f38715c8-0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")
	//req.Header.Set("X-User-ID", "40581705")
	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)

		return s, err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println(string(data))
		fmt.Println(err)
		return s, err
	}
	checktag(&s, exclude)
	if len(s.Body.IllustManga.Data) < 1 {
		return s, err
	}
	return s, nil
}
func getIllust(id string) (Illust, error) {
	var l Illust
	fmt.Println("传入id：" + id)
	baseURL := "https://www.pixiv.net/ajax/illust/" + id + "?"
	q := url.Values{}
	q.Add("lang", "zh")
	q.Add("version", "1514cd4826094c32a804b4de6def5f2209963922")

	// 拼接 URL
	url := baseURL + q.Encode()

	// 创建一个新的请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return l, nil
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Baggage", "sentry-environment=production,sentry-release=1514cd4826094c32a804b4de6def5f2209963922,sentry-public_key=7b15ebdd9cf64efb88cfab93783df02a,sentry-trace_id=34b30c0222da402190dab7f95ab1ca69,sentry-sample_rate=0.0001")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Cookie", "first_visit_datetime_pc=2024-10-27%2021%3A31%3A40; p_ab_id=2; p_ab_id_2=4; p_ab_d_id=692958056; yuid_b=QHl0dBI; c_type=25; privacy_policy_notification=0; a_type=0; b_type=1; PHPSESSID=40581705_rQmgRA1guZzvkWGVwysacRwWo2CcGh2C; device_token=6a83fbbbdb3dc55a762744511e49b9ca; _ga_MZ1NL4PHH0=GS1.1.1736429551.2.0.1736429573.0.0.0; privacy_policy_agreement=0; login_ever=yes; _gcl_au=1.1.987029509.1737216645; _gid=GA1.2.1484901853.1738498405; __cf_bm=wMi5L6y00VipuLNbBY4G3Ja9g2f6biFursX0W7PuHkc-1738505887-1.0.1.1-_xQv41W7vDvhF05MPgEFnO7ItcAQ1OugaUdT.AvsAHrszai3AgG5kY1Dw64jwjgUG7NvrcafLaTTMTGARgmrPLbDBsXaYtGvKzFo3wj.MVY; cf_clearance=AvLy6nioY1GkLeyqwVqEsTXh.a1IuxMKDRhK5bLXiqg-1738506871-1.2.1.1-O69GPQCbokRrCNUjwW.bSqjjdomkxW.e6ilk_BiRv__BXHX2lZPDwIM9py0OvU9Dy3Er5bDbS0385oZWusSUCmp4zh8AZzAeS2BsYKrTp0xduZn_phVENKu1lgfUimi2J3MCgCxTstdfMn_ESkWGwENXSor_Y_BQOLJ594BWls1rWo4aXvlxhQP2lmIzrCfUQD_C_sUONQCNGt.LiHPR8jAmNthvcqnusCY9y29bJaSnjf1gx_I3NS43vm1ARv1Is_pM4lmQ_pXt4HFlT8VjU8p1uIQoK4qrI4igi_OOfWQ; _gat_UA-1830249-3=1; _ga=GA1.1.1094039263.1730032302; _ga_75BBYNYN9J=GS1.1.1738505888.4.1.1738506875.0.0.0")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Referer", "https://www.pixiv.net/artworks/69956142")
	req.Header.Set("Sec-CH-UA", `"Not A(Brand";v="8", "Chromium";v="132", "Microsoft Edge";v="132"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sentry-Trace", "34b30c0222da402190dab7f95ab1ca69-a5e37ce88c2f6f31-0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/132.0.0.0 Safari/537.36 Edg/132.0.0.0")
	req.Header.Set("X-User-ID", "40581705")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return l, err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return l, err
	}
	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println(string(body))
		return l, err
	}
	// 打印响应
	return l, nil
}
func checktag(a *Search, b []string) {
	tagSet := make(map[string]struct{})
	for _, v := range b {
		tagSet[v] = struct{}{}
	}
	// 收集要删除的元素
	var indicesToRemove []int
	for k1, v1 := range a.Body.IllustManga.Data {
		for _, v2 := range v1.Tags {
			if _, exists := tagSet[v2]; exists {
				indicesToRemove = append(indicesToRemove, k1)
				break // 找到一个匹配后，跳出内层循环
			}
		}
	}
	// 删除元素
	for i := len(indicesToRemove) - 1; i >= 0; i-- {
		index := indicesToRemove[i]
		a.Body.IllustManga.Data = append(a.Body.IllustManga.Data[:index], a.Body.IllustManga.Data[index+1:]...)
	}
}

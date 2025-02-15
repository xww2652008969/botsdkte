package ff14logs

import (
	"fmt"
	"github.com/nfnt/resize"
	"github.com/xww2652008969/wbot/client/utils"
	"image"
	"image/jpeg"
	"image/png"
)

type Zone struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Difficulty int    `json:"difficulty"`
}
type Ouatdata struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}
type Token struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type GraphqlQuery struct {
	Query string `json:"query"`
}
type FF14Logdata struct {
	Data struct {
		CharacterData struct {
			Character struct {
				ZoneRankings struct {
					BestPerformanceAverage   float64 `json:"bestPerformanceAverage"`
					MedianPerformanceAverage float64 `json:"medianPerformanceAverage"`
					Difficulty               int     `json:"difficulty"`
					Metric                   string  `json:"metric"`
					Partition                int     `json:"partition"`
					Zone                     int     `json:"zone"`
					AllStars                 []struct {
						Partition      int     `json:"partition"`
						Spec           string  `json:"spec"`
						Points         float64 `json:"points"`
						PossiblePoints int     `json:"possiblePoints"`
						Rank           int     `json:"rank"`
						RegionRank     int     `json:"regionRank"`
						ServerRank     int     `json:"serverRank"`
						RankPercent    float64 `json:"rankPercent"`
						Total          int     `json:"total"`
					} `json:"allStars"`
					Rankings []struct {
						Encounter struct {
							Id   int    `json:"id"`
							Name string `json:"name"`
						} `json:"encounter"`
						RankPercent   float64 `json:"rankPercent"`
						MedianPercent float64 `json:"medianPercent"`
						LockedIn      bool    `json:"lockedIn"`
						TotalKills    int     `json:"totalKills"`
						FastestKill   int     `json:"fastestKill"`
						AllStars      struct {
							Points         float64 `json:"points"`
							PossiblePoints int     `json:"possiblePoints"`
							Partition      int     `json:"partition"`
							Rank           int     `json:"rank"`
							RegionRank     int     `json:"regionRank"`
							ServerRank     int     `json:"serverRank"`
							RankPercent    float64 `json:"rankPercent"`
							Total          int     `json:"total"`
						} `json:"allStars"`
						Spec       string  `json:"spec"`
						BestSpec   string  `json:"bestSpec"`
						BestAmount float64 `json:"bestAmount"`
					} `json:"rankings"`
				} `json:"zoneRankings"`
			} `json:"character"`
		} `json:"characterData"`
	} `json:"data"`
}
type RankingInfo struct {
	encounterId   int
	encounterName string
	rankPercent   float64
	totalKills    int
	spec          string
	bestAmount    float64
}
type Bossicon struct {
	bossicon map[int]image.Image
	jobicon  map[string]image.Image
}

func (b *Bossicon) geticon(id int) image.Image {
	if _, ok := b.bossicon[id]; ok {
		return b.bossicon[id]
	}
	res, err := utils.Httpget(fmt.Sprintf("https://assets.rpglogs.com/img/ff/bosses/%d-icon.jpg?v=2", id), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	j, err := jpeg.Decode(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	j = resize.Resize(64, 64, j, resize.Lanczos3)
	b.bossicon[id] = j
	return j
}
func (b *Bossicon) getjobicon(job string) image.Image {
	if _, ok := b.jobicon[job]; ok {
		return b.jobicon[job]
	}
	res, err := utils.Httpget(fmt.Sprintf("https://assets.rpglogs.com/img/ff/icons/%s.png", job), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	j, err := png.Decode(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	j = resize.Resize(32, 32, j, resize.Lanczos3)
	b.jobicon[job] = j
	return j
}

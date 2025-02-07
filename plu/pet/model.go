package pet

type key struct {
	Version     string `json:"version"`
	ApiVersion  int    `json:"api_version"`
	GraphicsApi string `json:"graphics_api"`
	Templates   []struct {
		Id       string `json:"id"`
		Metadata struct {
			ApiVersion            int      `json:"apiVersion"`
			TemplateVersion       int      `json:"templateVersion"`
			Alias                 []string `json:"alias"`
			Tags                  []string `json:"tags"`
			Author                string   `json:"author"`
			Desc                  string   `json:"desc"`
			Hidden                bool     `json:"hidden"`
			InRandomList          bool     `json:"inRandomList"`
			DefaultTemplateWeight int      `json:"defaultTemplateWeight"`
			Preview               *string  `json:"preview"`
		} `json:"metadata"`
	} `json:"templates"`
}
type reqstruct struct {
	Id    string `json:"id"`
	Image struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"image"`
	Text struct {
		To string `json:"to"`
	} `json:"text"`
}

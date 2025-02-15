package openaiutil

type OpenAi struct {
	Url         string
	Header      map[string]string
	RequestBody RequestBody
}

type MessageStruct struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type RequestBody struct {
	Model       string          `json:"model"`
	Messages    []MessageStruct `json:"messages"`
	MaxTokens   int             `json:"max_tokens"`
	Temperature float32         `json:"temperature"`
	Stream      bool            `json:"stream"`
}

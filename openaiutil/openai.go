package openaiutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Newclient(apiurl, apikey, model string, maxtokens int, prompt string) *OpenAi {
	var openai OpenAi
	openai.Url = apiurl
	openai.Header = make(map[string]string)
	openai.Header["Authorization"] = "Bearer " + apikey
	openai.Header["Content-Type"] = "application/json"
	var requestBody RequestBody
	requestBody.Model = model
	requestBody.MaxTokens = maxtokens
	requestBody.Temperature = 0.6
	requestBody.Stream = false
	var m MessageStruct
	m.Role = "system"
	m.Content = prompt
	requestBody.Messages = make([]MessageStruct, 0)
	requestBody.Messages = append(requestBody.Messages, m)
	openai.RequestBody = requestBody
	return &openai
}
func (a OpenAi) search(keywored string) {
	var m MessageStruct
	m.Role = "user"
	m.Content = keywored
	a.RequestBody.Messages = append(a.RequestBody.Messages, m)
	a.RequestBody.Messages = append(a.RequestBody.Messages, m)
	data, _ := json.Marshal(a.RequestBody)
	req, _ := http.NewRequest("POST", a.Url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", a.Header["Content-Type"])
	req.Header.Set("Authorization", a.Header["Authorization"])
	fmt.Println(req)
	res, _ := http.DefaultClient.Do(req)
	d, _ := io.ReadAll(res.Body)
	fmt.Println(string(d))
}

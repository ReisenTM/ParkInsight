package ai_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// DeepSeek源
const Url = "https://api.deepseek.com/chat/completions"

func DSRequest(r Request) (res *http.Response, err error) {
	method := "POST"
	byteData, _ := json.Marshal(r)
	req, err := http.NewRequest(method, Url, bytes.NewBuffer(byteData))
	if err != nil {
		logrus.Errorf("请求参数失败 %s", err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "sk-c616351b2e194417b07b3b870d775e0d"))
	req.Header.Add("Content-Type", "application/json")

	res, err = http.DefaultClient.Do(req)
	return
}
func DSToChat(content string) (msg string, err error) {
	r := Request{
		Model: "deepseek-chat",
		Messages: []Message{
			{
				Role:    "system",
				Content: chatPrompt,
			},
			{
				Role:    "user",
				Content: content,
			},
		},
	}
	res, err := DSRequest(r)
	if err != nil {
		return
	}
	body, _ := io.ReadAll(res.Body)
	var response ChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		logrus.Errorf("解析失败 %s %s", err, string(body))
		return
	}
	if len(response.Choices) > 0 {
		msg = response.Choices[0].Message.Content
		return
	}
	logrus.Errorf("未获取到数据 %s ", string(body))
	return
}

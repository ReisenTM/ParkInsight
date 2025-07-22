package ai_service

import (
	_ "embed"
)

// ChatGPT源
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ChatResponse struct {
	Id      string `json:"id"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		PromptTokens            int `json:"prompt_tokens"`
		CompletionTokens        int `json:"completion_tokens"`
		TotalTokens             int `json:"total_tokens"`
		CompletionTokensDetails struct {
			AudioTokens     int `json:"audio_tokens"`
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"completion_tokens_details"`
		PromptTokensDetails struct {
			AudioTokens  int `json:"audio_tokens"`
			CachedTokens int `json:"cached_tokens"`
		} `json:"prompt_tokens_details"`
	} `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

//go:embed chat.prompt
var chatPrompt string

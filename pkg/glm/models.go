package glm

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CallGLMResponse struct {
	Created   int    `json:"created"`
	Id        string `json:"id"`
	Model     string `json:"model"`
	RequestId string `json:"request_id"`
	Choices   []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

package main

type Session struct {
    AlreadyExists bool
    Id string
    Content string
}

type UserFile struct {
    Name string
    Contents string
}

type Model struct {
    Id string `json:"id"`
    Object string `json:"object"`
}

type ModelsRequest struct {
    Object string `json:"object"`
    Data []Model `json:"data"`
}

type Payload struct {
    Model string `json:"model"`
    Messages []Message `json:"messages"`
    Temperature float32 `json:"temperature"`
}

type Message struct {
	Role   string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int   `json:"index"`
	FinishReason string `json:"finish_reason"`
	Message      Message `json:"message"`
}

type Usage struct {
	PromptTokens int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens   int `json:"total_tokens"`
}

type Data struct {
	Created    int     `json:"created"`
	Object     string  `json:"object"`
	ID         string  `json:"id"`
	Model      string  `json:"model"`
	Choices    []Choice `json:"choices"`
	Usage      Usage   `json:"usage"`
}


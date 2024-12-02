package main

type OpenAIRequest struct {
	// A developer-provided per-request id that will be used to match outputs to inputs. Must be unique for each request in a batch.
	CustomID string             `json:"custom_id"`
	Method   string             `json:"method"`
	URL      string             `json:"url"`
	Body     ChatCompletionBody `json:"body"`
}

type ChatCompletionBody struct {
	Model            string         `json:"model"`
	Messages         []Messages     `json:"messages"`
	ResponseFormat   ResponseFormat `json:"response_format"`
	Temperature      int            `json:"temperature"`
	MaxTokens        int            `json:"max_tokens"`
	TopP             int            `json:"top_p"`
	FrequencyPenalty int            `json:"frequency_penalty"`
	PresencePenalty  int            `json:"presence_penalty"`
}
type Content struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
type Messages struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}
type Items struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Words struct {
	Type        string `json:"type"`
	Items       Items  `json:"items"`
	Description string `json:"description"`
}
type Definition struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
type Properties struct {
	Words      Words      `json:"words"`
	Definition Definition `json:"definition"`
}
type Schema struct {
	Type                 string     `json:"type"`
	Required             []string   `json:"required"`
	Properties           Properties `json:"properties"`
	AdditionalProperties bool       `json:"additionalProperties"`
}
type JSONSchema struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
	Strict bool   `json:"strict"`
}
type ResponseFormat struct {
	Type       string     `json:"type"`
	JSONSchema JSONSchema `json:"json_schema"`
}

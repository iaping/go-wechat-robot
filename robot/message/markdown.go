package message

type Markdown struct {
	Message
	Body *MarkdownBody `json:"markdown"`
}

func NewMarkdownSimple(content string) *Markdown {
	body := &MarkdownBody{
		Content: content,
	}

	return NewMarkdown(body)
}

func NewMarkdown(body *MarkdownBody) *Markdown {
	message := &Markdown{
		Body: body,
	}

	message.Type = "markdown"

	return message
}

type MarkdownBody struct {
	Content string `json:"content"`
}

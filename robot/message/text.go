package message

type Text struct {
	Message
	Body *TextBody `json:"text"`
}

func NewTextSimple(content string, isMentionedAll bool) *Text {
	body := &TextBody{
		Content: content,
	}

	if isMentionedAll {
		body.SetMentionedAll()
	}

	return NewText(body)
}

func NewText(body *TextBody) *Text {
	message := &Text{
		Body: body,
	}

	message.Type = "text"

	return message
}

type TextBody struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

func (b *TextBody) SetMentionedAll() {
	b.MentionedMobileList = []string{"@all"}
}

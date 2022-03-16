package message

type IMessage interface {
	GetType() string
}

type Message struct {
	Type string `json:"msgtype"`
}

func (m *Message) GetType() string {
	return m.Type
}

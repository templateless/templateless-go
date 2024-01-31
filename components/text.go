package components

type Text struct {
	Id   ComponentId `json:"id"`
	Text string      `json:"text"`
}

func NewText(text string) *Text {
	return &Text{
		Id:   ComponentIdText,
		Text: text,
	}
}

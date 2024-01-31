package components

type Button struct {
	Id   ComponentId `json:"id"`
	Text string      `json:"text"`
	URL  string      `json:"url"`
}

func NewButton(text, url string) *Button {
	return &Button{
		Id:   ComponentIdButton,
		Text: text,
		URL:  url,
	}
}

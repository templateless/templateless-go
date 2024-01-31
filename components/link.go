package components

type Link struct {
	Id   ComponentId `json:"id"`
	Text string      `json:"text"`
	URL  string      `json:"url"`
}

func NewLink(text, url string) *Link {
	return &Link{
		Id:   ComponentIdLink,
		Text: text,
		URL:  url,
	}
}

package components

type ViewInBrowser struct {
	Id   ComponentId `json:"id"`
	Text string      `json:"text"`
}

func NewViewInBrowser(text string) *ViewInBrowser {
	return &ViewInBrowser{
		Id:   ComponentIdViewInBrowser,
		Text: text,
	}
}

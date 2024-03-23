package components

type Image struct {
	Id     ComponentId `json:"id"`
	Src    string      `json:"src"`
	Alt    *string     `json:"alt"`
	Width  *int        `json:"width"`
	Height *int        `json:"height"`
	URL    *string     `json:"url"`
}

func NewImage(src string, url *string, width, height *int, alt *string) *Image {
	return &Image{
		Id:     ComponentIdImage,
		Src:    src,
		Alt:    alt,
		Width:  width,
		Height: height,
		URL:    url,
	}
}

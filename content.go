package templateless

import (
	"errors"

	"github.com/templateless/templateless-go/components"
)

type Theme string

const (
	Unstyled Theme = "UNSTYLED"
	Simple   Theme = "SIMPLE"
)

type Content struct {
	Version_ uint8                    `json:"version"`
	Theme_   Theme                    `json:"theme"`
	Header_  []components.Component   `json:"header"`
	Body_    [][]components.Component `json:"body"`
	Footer_  []components.Component   `json:"footer"`
}

func NewContent() *Content {
	return &Content{
		Version_: 0,
		Theme_:   Unstyled,
		Header_:  []components.Component{},
		Body_:    [][]components.Component{{}},
		Footer_:  []components.Component{},
	}
}

func (c *Content) Theme(theme Theme) *Content {
	c.Theme_ = theme
	return c
}

func (c *Content) Header(header Collection) *Content {
	c.Header_ = header.Components
	return c
}

func (c *Content) Footer(footer Collection) *Content {
	c.Footer_ = footer.Components
	return c
}

func (c *Content) Button(text, url string) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewButton(text, url))
	return c
}

func (c *Content) Image(src string) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewImage(src, nil, nil, nil, nil))
	return c
}

func (c *Content) Link(text, url string) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewLink(text, url))
	return c
}

func (c *Content) Otp(text string) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewOtp(text))
	return c
}

func (c *Content) Socials(data []components.SocialItem) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewSocials(data))
	return c
}

func (c *Content) Text(text string) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewText(text))
	return c
}

func (c *Content) ViewInBrowser() *Content {
	c.Body_[0] = append(c.Body_[0], components.NewViewInBrowser(nil))
	return c
}

func (c *Content) QrCode(url string) *Content {
	c.Body_[0] = append(c.Body_[0], components.QrCodeURL(url))
	return c
}

func (c *Content) StoreBadges(data []components.StoreBadgeItem) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewStoreBadges(data))
	return c
}

func (c *Content) Signature(text string) *Content {
	c.Body_[0] = append(c.Body_[0], components.NewSignature(text, nil))
	return c
}

func (c *Content) Component(component components.Component) *Content {
	c.Body_[0] = append(c.Body_[0], component)
	return c
}

func (c *Content) Build() (*Content, error) {
	if c == nil {
		return nil, errors.New("Content is nil")
	}

	newContent := *c
	return &newContent, nil
}

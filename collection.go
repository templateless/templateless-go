package templateless

import (
	"errors"

	"github.com/templateless/templateless-go/components"
)

type Collection struct {
	Components []components.Component `json:"components"`
}

func NewHeader() *Collection {
	return &Collection{
		Components: make([]components.Component, 0),
	}
}

func NewFooter() *Collection {
	return &Collection{
		Components: make([]components.Component, 0),
	}
}

func (c *Collection) Button(text, url string) *Collection {
	c.Components = append(c.Components, components.NewButton(text, url))
	return c
}

func (c *Collection) Image(src string) *Collection {
	c.Components = append(c.Components, components.NewImage(src, nil, nil, nil, nil))
	return c
}

func (c *Collection) Link(text, url string) *Collection {
	c.Components = append(c.Components, components.NewLink(text, url))
	return c
}

func (c *Collection) Otp(text string) *Collection {
	c.Components = append(c.Components, components.NewOtp(text))
	return c
}

func (c *Collection) Socials(data []components.SocialItem) *Collection {
	c.Components = append(c.Components, components.NewSocials(data))
	return c
}

func (c *Collection) Text(text string) *Collection {
	c.Components = append(c.Components, components.NewText(text))
	return c
}

func (c *Collection) ViewInBrowser() *Collection {
	c.Components = append(c.Components, components.NewViewInBrowser(nil))
	return c
}

func (c *Collection) QrCode(url string) *Collection {
	c.Components = append(c.Components, components.QrCodeURL(url))
	return c
}

func (c *Collection) StoreBadges(data []components.StoreBadgeItem) *Collection {
	c.Components = append(c.Components, components.NewStoreBadges(data))
	return c
}

func (c *Collection) Signature(text string) *Collection {
	c.Components = append(c.Components, components.NewSignature(text, nil))
	return c
}

func (c *Collection) Component(component components.Component) *Collection {
	c.Components = append(c.Components, component)
	return c
}

func (c *Collection) Build() (*Collection, error) {
	if c == nil {
		return nil, errors.New("Content is nil")
	}

	newCollection := *c
	return &newCollection, nil
}

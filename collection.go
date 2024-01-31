package templateless

import (
	"errors"

	"github.com/templateless/templateless-go/components"
)

type Collection struct {
	Components []components.Component `json:"components"`
}

func NewCollection() *Collection {
	return &Collection{
		Components: make([]components.Component, 0),
	}
}

func (c *Collection) Button(text, url string) *Collection {
	c.Components = append(c.Components, components.NewButton(text, url))
	return c
}

func (c *Collection) Image(src string, url string, width, height int, alt string) *Collection {
	c.Components = append(c.Components, components.NewImage(src, alt, url, width, height))
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

func (c *Collection) ViewInBrowser(text string) *Collection {
	c.Components = append(c.Components, components.NewViewInBrowser(text))
	return c
}

func (c *Collection) Build() (*Collection, error) {
	if c == nil {
		return nil, errors.New("Content is nil")
	}

	newCollection := *c
	return &newCollection, nil
}

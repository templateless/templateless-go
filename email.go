package templateless

import "errors"

type EmailOptions struct {
	DeleteAfter *uint64 `json:"deleteAfter,omitempty"`
}

type Email struct {
	to      map[EmailAddress]struct{}
	subject string
	content Content
	options EmailOptions
}

type EmailSlice struct {
	To      []EmailAddress `json:"to"`
	Subject string         `json:"subject"`
	Content Content        `json:"content"`
	Options EmailOptions   `json:"options"`
}

func NewEmail() *Email {
	content := NewContent()
	return &Email{
		to:      make(map[EmailAddress]struct{}),
		subject: "",
		content: *content,
		options: EmailOptions{DeleteAfter: nil},
	}
}

func (e *Email) To(emailAddress EmailAddress) *Email {
	e.to[emailAddress] = struct{}{}
	return e
}

func (e *Email) ToMany(emailAddresses []EmailAddress) *Email {
	for _, emailAddress := range emailAddresses {
		e.to[emailAddress] = struct{}{}
	}
	return e
}

func (e *Email) Subject(subject string) *Email {
	e.subject = subject
	return e
}

func (e *Email) Content(content Content) *Email {
	e.content = content
	return e
}

func (e *Email) DeleteAfter(seconds uint64) *Email {
	e.options.DeleteAfter = &seconds
	return e
}

func (e *Email) Build() (*Email, error) {
	if e == nil {
		return nil, errors.New("Content is nil")
	}

	newEmail := *e
	return &newEmail, nil
}

func (e *Email) ToSlice() *EmailSlice {
	var to []EmailAddress
	for email := range e.to {
		to = append(to, email)
	}

	email := &EmailSlice{
		To:      to,
		Subject: e.subject,
		Content: e.content,
		Options: e.options,
	}

	return email
}

package templateless

type EmailAddress struct {
	Name  *string `json:"name,omitempty"`
	Email string  `json:"email"`
}

func NewEmailAddress(email string) *EmailAddress {
	return &EmailAddress{
		Email: email,
	}
}

func NewEmailAddressWithName(name, email string) *EmailAddress {
	return &EmailAddress{
		Name:  &name,
		Email: email,
	}
}

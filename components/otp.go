package components

type Otp struct {
	Id   ComponentId `json:"id"`
	Text string      `json:"text"`
}

func NewOtp(text string) *Otp {
	return &Otp{
		Id:   ComponentIdOtp,
		Text: text,
	}
}

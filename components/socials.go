package components

type Service string

const (
	Website Service = "WEBSITE"
	Email   Service = "EMAIL"
	Twitter Service = "TWITTER"
	X       Service = "X"
	Github  Service = "GITHUB"
)

type SocialItem struct {
	Service Service `json:"service"`
	Value   string  `json:"value"`
}

func NewSocialItem(service Service, value string) *SocialItem {
	return &SocialItem{
		Service: service,
		Value:   value,
	}
}

type Socials struct {
	ID   ComponentId  `json:"id"`
	Data []SocialItem `json:"data"`
}

func NewSocials(data []SocialItem) *Socials {
	return &Socials{
		ID:   ComponentIdSocials,
		Data: data,
	}
}

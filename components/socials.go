package components

type Service string

const (
	Website   Service = "WEBSITE"
	Email     Service = "EMAIL"
	Phone     Service = "PHONE"
	Facebook  Service = "FACEBOOK"
	YouTube   Service = "YOUTUBE"
	Twitter   Service = "TWITTER"
	X         Service = "X"
	GitHub    Service = "GITHUB"
	Instagram Service = "INSTAGRAM"
	LinkedIn  Service = "LINKEDIN"
	Slack     Service = "SLACK"
	Discord   Service = "DISCORD"
	TikTok    Service = "TIKTOK"
	Snapchat  Service = "SNAPCHAT"
	Threads   Service = "THREADS"
	Telegram  Service = "TELEGRAM"
	Mastodon  Service = "MASTODON"
	Rss       Service = "RSS"
)

type SocialItem struct {
	Service Service `json:"key"`
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

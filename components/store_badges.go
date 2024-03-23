package components

type StoreBadge string

const (
	AppStore       StoreBadge = "APP_STORE"
	GooglePlay     StoreBadge = "GOOGLE_PLAY"
	MicrosoftStore StoreBadge = "MICROSOFT_STORE"
)

type StoreBadgeItem struct {
	Key   StoreBadge `json:"key"`
	Value string     `json:"value"`
}

func NewStoreBadgeItem(key StoreBadge, value string) *StoreBadgeItem {
	return &StoreBadgeItem{
		Key:   key,
		Value: value,
	}
}

type StoreBadges struct {
	ID   ComponentId      `json:"id"`
	Data []StoreBadgeItem `json:"data"`
}

func NewStoreBadges(data []StoreBadgeItem) *StoreBadges {
	return &StoreBadges{
		ID:   ComponentIdStoreBadges,
		Data: data,
	}
}

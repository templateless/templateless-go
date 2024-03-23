package components

type SignatureFont string

const (
	ReenieBeanie SignatureFont = "REENIE_BEANIE"
	MeowScript   SignatureFont = "MEOW_SCRIPT"
	Caveat       SignatureFont = "CAVEAT"
	Zeyada       SignatureFont = "ZEYADA"
	Petemoss     SignatureFont = "PETEMOSS"
)

type Signature struct {
	Id   ComponentId    `json:"id"`
	Text string         `json:"text"`
	Font *SignatureFont `json:"font"`
}

func NewSignature(text string, font *SignatureFont) *Signature {
	return &Signature{
		Id:   ComponentIdSignature,
		Text: text,
		Font: font,
	}
}

package components

type ComponentId string

const (
	ComponentIdButton        ComponentId = "BUTTON"
	ComponentIdImage         ComponentId = "IMAGE"
	ComponentIdLink          ComponentId = "LINK"
	ComponentIdOtp           ComponentId = "OTP"
	ComponentIdPoweredBy     ComponentId = "POWERED_BY"
	ComponentIdSocials       ComponentId = "SOCIALS"
	ComponentIdText          ComponentId = "TEXT"
	ComponentIdViewInBrowser ComponentId = "VIEW_IN_BROWSER"
)

type Component interface {
}

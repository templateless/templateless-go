package main

import (
	"log"
	"os"

	"github.com/templateless/templateless-go"
	"github.com/templateless/templateless-go/components"
)

func main() {
	apiKey := os.Getenv("TEMPLATELESS_API_KEY")
	if apiKey == "" {
		log.Fatal("Set TEMPLATELESS_API_KEY to your Templateless API key")
	}

	emailAddress := os.Getenv("TEMPLATELESS_EMAIL_ADDRESS")
	if emailAddress == "" {
		log.Fatal("Set TEMPLATELESS_EMAIL_ADDRESS to your own email address")
	}

	width := 100
	alt := "MyCo"
	header, _ := templateless.NewHeader().
		Component(
			components.NewImage(
				"https://templateless.net/myco.webp",
				nil,
				&width,
				nil,
				&alt,
			),
		).
		Build()

	footer, _ := templateless.NewFooter().
		Text("If you did not sign up for a MyCo account, please ignore this email.\nThis link will expire in 24 hours.").
		Socials([]components.SocialItem{
			*components.NewSocialItem(components.Twitter, "MyCo"),
			*components.NewSocialItem(components.GitHub, "MyCo"),
		}).
		Link("Unsubscribe", "https://example.com").
		Build()

	verifyEmailLink := "https://example.com/verify?token=ABC"

	content, _ := templateless.NewContent().
		Theme(templateless.Simple).
		Header(*header).
		Text("Hi there,").
		Text("Welcome to **MyCo**! We're excited to have you on board. Before we get started, we need to verify your email address.").
		Text("Please confirm your email by clicking the button below:").
		Button("Verify Email", verifyEmailLink).
		Text("Or use the link below:").
		Link(verifyEmailLink, verifyEmailLink).
		Footer(*footer).
		Build()

	email, _ := templateless.NewEmail().
		To(*templateless.NewEmailAddress(emailAddress)).
		Subject("Confirm your email").
		Content(*content).
		Build()

	templateless.NewTemplateless(apiKey).
		Send(*email)
}

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
		Socials([]components.SocialItem{
			*components.NewSocialItem(components.Twitter, "MyCo"),
			*components.NewSocialItem(components.GitHub, "MyCo"),
		}).
		Build()

	content, _ := templateless.NewContent().
		Theme(templateless.Simple).
		Header(*header).
		Text("Hey Alex,").
		Text("I'm Jamie, founder of **MyCo**.").
		Text("Could you spare a moment to reply to this email with your thoughts on our service? Your feedback is invaluable and directly influences our improvements.").
		Text("When you hit reply, your email will go directly to me, and I read each and every one.").
		Text("Thanks for your support,").
		Signature("Jamie Parker").
		Text("Jamie Parker\n\nFounder @ [MyCo](https://example.com)").
		Footer(*footer).
		Build()

	email, _ := templateless.NewEmail().
		To(*templateless.NewEmailAddress(emailAddress)).
		Subject("Thoughts on service?").
		Content(*content).
		Build()

	templateless.NewTemplateless(apiKey).
		Send(*email)
}

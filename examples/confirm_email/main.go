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

	header, _ := templateless.NewHeader().
		Text("# ExampleApp").
		Build()

	footer, _ := templateless.NewFooter().
		Socials([]components.SocialItem{
			*components.NewSocialItem(components.Twitter, "Username"),
			*components.NewSocialItem(components.GitHub, "Username"),
		}).
		Link("Unsubscribe", "https://example.com/unsubscribe?id=123").
		Build()

	content, _ := templateless.NewContent().
		Header(*header).
		Text("Hello world").
		Footer(*footer).
		Build()

	email, _ := templateless.NewEmail().
		To(*templateless.NewEmailAddress(emailAddress)).
		Subject("Confirm your email").
		Content(*content).
		Build()

	result, _ := templateless.NewTemplateless(apiKey).
		Send(*email)

	log.Println(result)
}

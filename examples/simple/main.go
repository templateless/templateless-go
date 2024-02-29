package main

import (
	"log"
	"os"

	"github.com/templateless/templateless-go"
)

func main() {
	apiKey := os.Getenv("TEMPLATELESS_API_KEY")
	if apiKey == "" {
		log.Println("Set TEMPLATELESS_API_KEY to your Templateless API key")
		return
	}

	emailAddress := os.Getenv("TEMPLATELESS_EMAIL_ADDRESS")
	if emailAddress == "" {
		log.Println("Set TEMPLATELESS_EMAIL_ADDRESS to your own email address")
		return
	}

	content, _ := templateless.NewContent().
		Text("Hello world").
		Build()

	email, _ := templateless.NewEmail().
		To(*templateless.NewEmailAddress(emailAddress)).
		Subject("Hello 👋").
		Content(*content).
		Build()

	result, _ := templateless.NewTemplateless(apiKey).
		Send(*email)

	log.Println(result)
}

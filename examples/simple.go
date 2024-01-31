package main

import (
	"log"

	"github.com/templateless/templateless-go"
)

func main() {
	content, _ := templateless.NewContent().
		Text("Hello world").
		Build()

	emailAddress := templateless.NewEmailAddress("user@example.com")
	email, _ := templateless.NewEmail().
		To(*emailAddress).
		Subject("Hello 👋").
		Content(*content).
		Build()

	result, _ := templateless.NewTemplateless("<YOUR_API_KEY>").
		Send(*email)

	log.Println(result)
}

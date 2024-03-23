package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

	appStoreLink := "https://apps.apple.com/us/app/example/id1234567890"
	googlePlayLink :=
		"https://play.google.com/store/apps/details?id=com.example"

	footer, _ := templateless.NewFooter().
		StoreBadges([]components.StoreBadgeItem{
			*components.NewStoreBadgeItem(components.AppStore, appStoreLink),
			*components.NewStoreBadgeItem(components.GooglePlay, googlePlayLink),
		}).
		Socials([]components.SocialItem{
			*components.NewSocialItem(components.Twitter, "MyCo"),
			*components.NewSocialItem(components.GitHub, "MyCo"),
		}).
		Build()

	content, _ := templateless.NewContent().
		Theme(templateless.Simple).
		Header(*header).
		Text("Hey Alex,").
		Text("Thank you for choosing MyCo! To get started with our mobile experience, simply pair your account with our mobile app.").
		Text("Here's how to do it:").
		Text(strings.Join([]string{
			fmt.Sprintf("1. Download the MyCo app from the [App Store](%s) or [Google Play](%s).", appStoreLink, googlePlayLink),
			"1. Open the app and select _Pair Device_.",
			"1. Scan the QR code below with your mobile device:",
		}, "\n")).
		QrCode("https://example.com/qr-code-link").
		Text("Enjoy your seamless experience across devices!").
		Footer(*footer).
		Build()

	email, _ := templateless.NewEmail().
		To(*templateless.NewEmailAddress(emailAddress)).
		Subject("How to Pair Device").
		Content(*content).
		Build()

	templateless.NewTemplateless(apiKey).
		Send(*email)
}

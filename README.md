<h1 align="center">
  <a href="https://templateless.com/">
    <img src="templateless.webp" alt="Templateless" width="450px">
  </a>
  <br />
</h1>

<p align="center">
  <b>Ship faster by treating email as code 🚀</b> <br />
</p>

<h4 align="center">
  <a href="https://templateless.com/">Website</a> &bull;
  <a href="https://app.templateless.com/">Get Your API Key</a> &bull;
  <a href="https://twitter.com/templateless">Twitter</a>
</h4>

---

[![GoDoc](https://godoc.org/github.com/templateless/templateless-go?status.svg)](https://godoc.org/github.com/templateless/templateless-go)
[![Github Actions](https://img.shields.io/github/actions/workflow/status/templateless/templateless-go/tests.yml)](https://github.com/templateless/templateless-go/actions)
[![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/Templateless)](https://twitter.com/templateless)

[Templateless](https://templateless.com) lets you generate and send transactional emails quickly and easily so you can focus on building your product.

It's perfect for SaaS, web apps, mobile apps, scripts and anywhere you have to send email programmatically.

## ✨ Features

- 👋 **Anti drag-and-drop by design** — emails are a part of your code
- ✅ **Components as code** — function calls turn into email HTML components
- 💻 **SDK for any language** — use your favorite [programming language](https://github.com/orgs/templateless/repositories)
- 🔍 **Meticulously tested** — let us worry about email client compatibility
- 💌 **Use your favorite ESP** — Amazon SES, SendGrid, Mailgun + more
- 💪 **Email infrastructure** — rate-limiting, retries, scheduling + more
- ⚡ **Batch sending** — send 1 email or 1,000 with one API call

## 🚀 Getting started

Use go get:

```bash
go get github.com/templateless/templateless-go@v0.1.0-alpha.4
```

Then import the package into your own code:

```go
import "github.com/templateless/templateless-go"
```

## 🔑 Get Your API Key

You'll need an API key for the example below ⬇️

[![Get Your API Key](https://img.shields.io/badge/Get_Your_API_Key-free-blue?style=for-the-badge)](https://app.templateless.com/)

- 3,000 emails per month
- All popular email provider integrations
- Start sending right away

## 👩‍💻 Quick example

This is all it takes to send a signup confirmation email:

```go
package main

import (
	"log"

	"github.com/templateless/templateless-go"
)

func main() {
	content, _ := templateless.NewContent().
		Text("Hi, please **confirm your email**:").
		Button("Confirm Email", "https://your-company.com/signup/confirm?token=XYZ").
		Build()

	emailAddress := templateless.NewEmailAddress("<YOUR_CUSTOMERS_EMAIL_ADDRESS>")
	email, _ := templateless.NewEmail().
		To(*emailAddress).
		Subject("Confirm your signup 👋").
		Content(*content).
		Build()

	templateless.NewTemplateless("<YOUR_API_KEY>").
		Send(*email)
}
```

There are more Go examples in the [examples](examples) folder ✨

> [!NOTE]
> 🚧 **The SDK is not stable yet.** This API might change as more features are added. Please watch the repo for the changes in the [CHANGELOG](CHANGELOG.md).

## 🏗 Debugging

You can generate _test API keys_ by activating the **Test Mode** in your dashboard. By using these keys, you'll be able to view your fully rendered emails without actually sending them.

When you use a test API key in your SDK, the following output will appear in your logs when you try to send an email:

```log
Templateless [TEST MODE]: Emailed user@example.com, preview: https://tmpl.sh/ATMxHLX4r9aE
```

The preview link will display the email, but you must be logged in to Templateless to view it.

## 🔳 Components

Emails are crafted programmatically by making function calls. There's no dealing with HTML or drag-and-drop builders.

All of the following components can be mixed and matched to create dynamic emails:

<details>
  <summary>Text / Markdown</summary>

Text component allow you to insert a paragraph. Each paragraph supports basic markdown:

- Bold text: `**bold text**`
- Italic text: `_italic text_`
- Link: `[link text](https://example.com)`
- Also a link: `<https://example.com>`
- Headers (h1-h6):

  - `# Big Header`
  - `###### Small Header`

- Unordered list:

  ```md
  - item one
  - item two
  - item three
  ```

- Ordered list:

  ```md
  1. item one
  1. item two
  1. item three
  ```

```go
templateless.NewContent().
	Text("## Thank you for signing up").
	Text("Please **verify your email** by [clicking here](https://example.com/confirm?token=XYZ)").
	Build()
```

</details>
<details><summary>Link</summary>

Link component adds an anchor tag. This is the same as a text component with the link written in markdown:

```go
templateless.NewContent().
	Link("Confirm Email", "https://example.com/confirm?token=XYZ").
	Build()
```

</details>
<details><summary>Button</summary>

Button can also be used as a call to action. Button color is set via your dashboard's app color.

```go
templateless.NewContent().
	Button("Confirm Email", "https://example.com/confirm?token=XYZ").
	Build()
```

</details>
<details><summary>Image</summary>

Image component will link to an image within your email. Keep in mind that a lot of email clients will prevent images from being loaded automatically for privacy reasons.

```go
// Simple
templateless.NewContent().
	Image("https://placekitten.com/300/200").
	Build()

// Clickable & with attributes
url := "https://example.com"
width := 300
height := 200
alt := "Alt text"
templateless.NewContent().
	Component(
		components.NewImage(
			"https://placekitten.com/300/200",
			&url,
			&width,
			&height,
			&alt,
		),
	).
	Build()
```

Only the `src` parameter is required; everything else is optional.

**If you have "Image Optimization" turned on:**

1. Your images will be cached and distributed by our CDN for faster loading. The cache does not expire. If you'd like to re-cache, simply append a query parameter to the end of your image url.
1. Images will be converted into formats that are widely supported by email clients. The following image formats will be processed automatically:

    - Jpeg
    - Png
    - Gif
    - WebP
    - Tiff
    - Ico
    - Bmp
    - Svg

1. Maximum image size is 5MB for free accounts and 20MB for paid accounts.
1. You can specify `width` and/or `height` if you'd like (they are optional). Keep in mind that images will be scaled down to fit within the email theme, if they're too large.

</details>
<details><summary>One-Time Password</summary>

OTP component is designed for showing temporary passwords and reset codes.

```go
templateless.NewContent().
	Text("Here's your **temporary login code**:").
	Otp("XY78-2BT0-YFNB-ALW9").
	Build()
```

</details>
<details><summary>Social Icons</summary>

You can easily add social icons with links by simply specifying the username. Usually, this component is placed in the footer of the email.

These are all the supported platforms:

```go
templateless.NewContent().
	Socials([]components.SocialItem{
		*components.NewSocialItem(components.Website, "https://example.com"),
		*components.NewSocialItem(components.Email, "username@example.com"),
		*components.NewSocialItem(components.Phone, "123-456-7890"), // `tel:` link
		*components.NewSocialItem(components.Facebook, "Username"),
		*components.NewSocialItem(components.YouTube, "ChannelID"),
		*components.NewSocialItem(components.Twitter, "Username"),
		*components.NewSocialItem(components.X, "Username"),
		*components.NewSocialItem(components.GitHub, "Username"),
		*components.NewSocialItem(components.Instagram, "Username"),
		*components.NewSocialItem(components.LinkedIn, "Username"),
		*components.NewSocialItem(components.Slack, "Org"),
		*components.NewSocialItem(components.Discord, "Username"),
		*components.NewSocialItem(components.TikTok, "Username"),
		*components.NewSocialItem(components.Snapchat, "Username"),
		*components.NewSocialItem(components.Threads, "Username"),
		*components.NewSocialItem(components.Telegram, "Username"),
		*components.NewSocialItem(components.Mastodon, "@Username@example.com"),
		*components.NewSocialItem(components.Rss, "https://example.com/blog"),
	}).
	Build()
```

</details>
<details><summary>View in Browser</summary>

If you'd like your recipients to be able to read the email in a browser, you can add the "view in browser" component that will automatically generate a link. Usually, this is placed in the header or footer of the email.

You can optionally provide the text for the link. If none is provided, default is used: "View in browser"

**Anyone who knows the link will be able to see the email.**

```go
templateless.NewContent().
	ViewInBrowser().
	Build()
```

</details>
<details><summary>Store Badges</summary>

Link to your mobile apps via store badges:

```go
templateless.NewContent().
	StoreBadges([]components.StoreBadgeItem{
		*components.NewStoreBadgeItem(components.AppStore, "https://apps.apple.com/us/app/example/id1234567890"),
		*components.NewStoreBadgeItem(components.GooglePlay, "https://play.google.com/store/apps/details?id=com.example"),
		*components.NewStoreBadgeItem(components.MicrosoftStore, "https://apps.microsoft.com/detail/example"),
	}).
	Build()
```

</details>
<details><summary>QR Code</summary>

You can also generate QR codes on the fly. They will be shown as images inside the email.

Here are all the supported data types:

```go
// URL
templateless.NewContent().
	QrCode("https://example.com").
	Build()

// Email
templateless.NewContent().
	Component(components.QrCodeEmail("user@example.com")).
	Build()

// Phone
templateless.NewContent().
	Component(components.QrCodePhone("123-456-7890")).
	Build()

// SMS / Text message
templateless.NewContent().
	Component(components.QrCodeSMS("123-456-7890")).
	Build()

// Geo coordinates
templateless.NewContent().
	Component(components.QrCodeCoordinates(37.773972, -122.431297)).
	Build()

// Crypto address (for now only Bitcoin and Ethereum are supported)
templateless.NewContent().
	Component(components.QrCodeCryptocurrencyAddress(components.Bitcoin, "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")).
	Build()

// You can also encode any binary data
templateless.NewContent().
	Component(components.NewQrCode([]byte{1, 2, 3})).
	Build()
```

</details>
<details><summary>Signature</summary>

Generated signatures can be added to your emails to give a bit of a personal touch. This will embed an image with your custom text using one of several available fonts:

```go
// Signature with a default font
templateless.NewContent().
	Signature("John Smith").
	Build()

// Signature with a custom font
templateless.NewContent().
	Component(components.NewSignature("John Smith", &components.ReenieBeanie)).
	Build()
```

These are the available fonts:

- `ReenieBeanie` [preview →](https://fonts.google.com/specimen/Reenie+Beanie)
- `MeowScript` [preview →](https://fonts.google.com/specimen/Meow+Script)
- `Caveat` [preview →](https://fonts.google.com/specimen/Caveat)
- `Zeyada` [preview →](https://fonts.google.com/specimen/Zeyada)
- `Petemoss` [preview →](https://fonts.google.com/specimen/Petemoss)

Signature should not exceed 64 characters. Only alphanumeric characters and most common symbols are allowed.

</details>


---

Components can be placed in the header, body and footer of the email. Header and footer styling is usually a bit different from the body (for example the text is smaller).

```go
header, _ := templateless.NewHeader(). // header of the email
  Text("Smaller text").
  Build()

content, _ := templateless.NewContent(). // body of the email
  Text("Normal text").
  Build()
```

Currently there are 2 themes to choose from: `Unstyled` and `Simple`

```go
content, _ := templateless.NewContent().
  Theme(templateless.Simple).
  Text("Hello world").
  Build()
```

## 🤝 Contributing

- Contributions are more than welcome
- Please **star this repo** for more visibility <3

## 📫 Get in touch

- For customer support feel free to email us at [github@templateless.com](mailto:github@templateless.com)

- Have suggestions or want to give feedback? Here's how to reach us:

    - For feature requests, please [start a discussion](https://github.com/templateless/templateless-go/discussions)
    - Found a bug? [Open an issue!](https://github.com/templateless/templateless-go/issues)
    - Say hi [@Templateless](https://twitter.com/templateless) 👋

## 🍻 License

[MIT](LICENSE)
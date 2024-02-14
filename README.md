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
go get github.com/templateless/templateless-go
```

Then import the package into your own code:

```go
import "github.com/templateless/templateless-go"
```

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

	result, _ := templateless.NewTemplateless("<YOUR_API_KEY>").
		Send(*email)

	log.Println(result)
}
```

> **Note**
> 🚧 **This SDK is not stable yet.** The API might change as more features are added. Please pay attention to the [CHANGELOG](CHANGELOG.md) for breaking changes.

Examples:

1. Get your **free API key** here: <https://app.templateless.com> ✨
1. There are more Go examples in the [examples](examples) folder

## 🤝 Contributing

- Contributions are more than welcome <3
- Please **star this repo** for more visibility ★

## 📫 Get in touch

- For customer support feel free to email us at [github@templateless.com](mailto:github@templateless.com)

- Have suggestions or want to give feedback? Here's how to reach us:

    - For feature requests, please [start a discussion](https://github.com/templateless/templateless-go/discussions)
    - Found a bug? [Open an issue!](https://github.com/templateless/templateless-go/issues)
    - We are also on Twitter: [@Templateless](https://twitter.com/templateless)

## 🍻 License

[MIT](LICENSE)
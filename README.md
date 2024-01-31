# Templateless Go

[![GoDoc](https://godoc.org/github.com/templateless/templateless-go?status.svg)](https://godoc.org/github.com/templateless/templateless-go)
[![Github Actions](https://img.shields.io/github/actions/workflow/status/templateless/templateless-go/tests.yml)](https://github.com/templateless/templateless-go/actions)

## What is this?

[Templateless](https://templateless.com) lets you quickly create and send emails with your favorite email provider without ever leaving your code editor.

Don't waste time messing around with HTML or HTML builders.

**Get your free API key [here](https://app.templateless.com).**

## Getting Started

Use go get.

```bash
go get github.com/templateless/templateless-go
```

Then import the package into your own code.

```go
import "github.com/templateless/templateless-go"
```

## Quick Example

```go
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
```

## License

[MIT](LICENSE)
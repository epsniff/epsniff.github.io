package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

/*
# Build app.wasm:
GOARCH=wasm GOOS=js go build -o /test-app/web/app.wasm

# Build and generate static website:
go run main.go
*/

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World from go-app!")
}

func main() {
	app.Route("/", func() app.Composer { return &hello{} })
	app.Route("/hello", func() app.Composer { return &hello{} })
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Resources:   app.GitHubPages("epsniff/epsniff.github.io"),
	})

	if err != nil {
		log.Fatal(err)
	}
}

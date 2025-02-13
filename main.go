package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

/*
GOARCH=wasm GOOS=js go build -o  ./web-app/app.wasm # Build app.wasm:
go run main.go # Build and generate static website:
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
	log.Default().Println("this is just a test log message.")
	return app.H1().Text("Hello World 222!")
}

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello2 struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello2) Render() app.UI {
	return app.H1().Text("Hey, lets talk!")
}

func main() {
	app.Route("/", func() app.Composer { return &hello{} })
	app.Route("/hello", func() app.Composer { return &hello2{} })
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

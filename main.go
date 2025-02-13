package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"github.com/epsniff/epsniff.github.io/src/components"
)

func main() {
	app.Route("/", func() app.Composer { return &components.Index{} })
	app.Route("/debug", func() app.Composer { return &components.Debug{} })
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Eric Sniff",
		Description: "Eric Sniff personal website",
		Resources:   app.GitHubPages("epsniff/epsniff.github.io"),
	})

	if err != nil {
		log.Fatal(err)
	}
}

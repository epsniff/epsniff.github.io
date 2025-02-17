package main

import (
	"log"

	"github.com/epsniff/epsniff.github.io/src/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	app.Route("/", func() app.Composer { return &components.Index{} })
	app.Route("/debug", func() app.Composer { return &components.Debug{} })
	app.Route("/line", func() app.Composer { return &components.LineCombo{} })
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

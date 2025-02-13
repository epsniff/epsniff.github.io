package main

import (
	"log"

	"github.com/maxence-charriere/go-app/v10/pkg/app"

	"github.com/epsniff/epsniff.github.io/src/components"
)

/*
GOARCH=wasm GOOS=js go build -o  ./web-app/app.wasm # Build app.wasm:
go run main.go # Build and generate static website
git add . && git commit -m "Update" && git push # Push to github
*/

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

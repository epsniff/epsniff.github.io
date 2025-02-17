package main

import (
	"log"

	"github.com/epsniff/epsniff.github.io/src/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	app.Route("/", func() app.Composer { return &components.Index{} })
	app.Route("/debug", func() app.Composer { return &components.Debug{} })
	app.Route("/line", func() app.Composer { return components.NewLineCombo().SetLineChart(components.NewLine()) })
	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Eric Sniff",
		Description: "Eric Sniff personal website",
		Resources:   app.GitHubPages("epsniff/epsniff.github.io"),
		Scripts: []string{
			"https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js",
			"https://go-echarts.github.io/go-echarts-assets/assets/themes/westeros.js",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Index struct {
	app.Compo
}

func (h *Index) Render() app.UI {
	return app.Div().
		Style("background-color", "#f0f0f0").
		Style("padding", "20px").
		Style("text-align", "center").
		Body(
			app.H1().Text("Go App Test"),
			app.P().Text("Welcome to my personal website!"),
			app.A().Href("/debug").Text("Debug"),
		)
}

package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Index struct {
	app.Compo
}

func (h *Index) Render() app.UI {
	return app.H1().Text("Hello World 42!")
}

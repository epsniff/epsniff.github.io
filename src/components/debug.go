package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Debug struct {
	app.Compo
}

func (h *Index) Debug() app.UI {
	return app.H1().Text("Debug")
}

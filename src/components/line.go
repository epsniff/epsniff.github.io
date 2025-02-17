package components

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"golang.org/x/exp/rand"
)

func NewLineCombo() *LineCombo {
	return &LineCombo{}
}

type LineCombo struct {
	app.Compo

	lineChart *charts.Line
	instance  app.Value
}

func (h *LineCombo) SetLineChart(l *charts.Line) *LineCombo {
	h.lineChart = l
	return h
}

func (a *LineCombo) Render() app.UI {
	return app.Div().Class("item").
		Style("width", a.lineChart.Initialization.Width).
		Style("height", a.lineChart.Initialization.Height)
}
func (a *LineCombo) OnDismount(ctx app.Context) {
	if a.instance != nil {
		a.instance.Call("dispose")
	}
}
func (a *LineCombo) OnMount(ctx app.Context) {
	var val string
	ctx.ObserveState("foo", &val).OnChange(
		func() {
			fmt.Println("onChange is called", val)
			if a.instance != nil {
				a.instance.Call("dispose")
			}
			n := NewLine()
			n.ID = a.lineChart.ID
			a.lineChart = n
			a.instance = app.Window().Get("echarts").
				Call("init", a.JSValue(), a.lineChart.Theme)
			ctx.Async(func() {
				options := app.Window().Get("JSON").Call("parse", a.lineChart.RenderSnippet().Option)
				a.instance.Call("setOption", options)
			})
		})
	ctx.Defer(func(context app.Context) {
		if a.instance != nil {
			a.instance.Call("dispose")
		}
		a.instance = app.Window().Get("echarts").
			Call("init", a.JSValue(), a.lineChart.Theme, `{ renderer : "canvas" }`)
		ctx.Async(func() {
			// jsonString, _ := json.Marshal(a.lineChart.JSON())
			// jsonStr := string(jsonString)
			options := app.Window().Get("JSON").Call("parse", a.lineChart.RenderSnippet().Option)
			a.instance.Call("setOption", options)
		})
	})
}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func NewLine() *charts.Line {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))
	return line
}

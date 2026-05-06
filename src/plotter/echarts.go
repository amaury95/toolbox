package plotter

import (
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Point struct {
	X float64
	Y float64
}

// PlotLine draws an interactive line chart in the default browser using go-echarts.
// It serves the page on 127.0.0.1, opens a tab/window, and blocks until the process is interrupted (Ctrl+C).
// Zoom on both axes: inside (scroll/drag on the plot), horizontal range slider on the bottom (x),
// vertical range slider on the right (y), toolbox data-zoom for area selection, and save-as-PNG.
func PlotLine(title, xLabel, yLabel string, points ...Point) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithXAxisOpts(opts.XAxis{
			Type: "value",
			Name: xLabel,
		}),
		charts.WithYAxisOpts(opts.YAxis{Name: yLabel}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show: opts.Bool(true),
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show: opts.Bool(true),
					Type: "png",
					Name: "chart",
				},
				DataZoom: &opts.ToolBoxFeatureDataZoom{
					Show: opts.Bool(true),
				},
			},
		}),
		// X + Y dataZoom: inside (wheel/drag) + axis-aligned sliders.
		charts.WithDataZoomOpts(
			opts.DataZoom{Type: "inside", XAxisIndex: 0},
			opts.DataZoom{Type: "inside", YAxisIndex: 0},
			opts.DataZoom{
				Type:       "slider",
				XAxisIndex: 0,
				Start:      0,
				End:        100,
			},
			opts.DataZoom{
				Type:         "slider",
				YAxisIndex:   0,
				Orient:       "vertical",
				Start:        0,
				End:          100,
			},
		),
	)

	items := make([]opts.LineData, 0, len(points))
	for _, p := range points {
		items = append(items, opts.LineData{
			Value: []interface{}{p.X, p.Y},
		})
	}

	line.AddSeries(yLabel, items,
		charts.WithLineChartOpts(opts.LineChart{
			Smooth: opts.Bool(true),
		}),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		line.Render(w)
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	url := "http://" + ln.Addr().String()

	go func() {
		time.Sleep(200 * time.Millisecond)
		_ = openBrowser(url)
	}()

	if err := http.Serve(ln, mux); err != nil {
		panic(err)
	}
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	return cmd.Start()
}

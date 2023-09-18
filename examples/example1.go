package main

import (
	"github.com/dimkauzh/vertex-engine/src/vertex"
	"github.com/dimkauzh/vertex-engine/src/vertex/draw"
	w "github.com/dimkauzh/vertex-engine/src/vertex/window"
)

func main() {
	vertex.Init()
	defer vertex.Quit()

	window := w.NewWindow(800, 600, "Vertex Engine")

	for w.Loop(window) {
		draw.SetBackgroundColor([3]float32{0.2, 0.2, 0.2})
		draw.DrawLine(-0.5, -0.5, 0.5, 0.5, [3]float32{1, 0, 0})
		draw.DrawRect(-0.5, 0.2, 0.5, 0.5, [3]float32{0, 1, 1})
		w.Refresh(window)
	}
}

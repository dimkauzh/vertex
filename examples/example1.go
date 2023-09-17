package main

import (
	"github.com/dimkauzh/vertex-engine/src/vertex"
	w "github.com/dimkauzh/vertex-engine/src/vertex/window"
)

func main() {
	vertex.Init()
	defer vertex.Quit()

	window := w.NewWindow(800, 600, "Vertex Engine")

	for w.Loop(window) {
		// Do something
	}
}

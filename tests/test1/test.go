package main

import (
	"fmt"

	"github.com/dimkauzh/vertex"
	"github.com/dimkauzh/vertex/src/draw"
	"github.com/dimkauzh/vertex/src/event"
	w "github.com/dimkauzh/vertex/src/window"
)

func main() {
	vertex.Init()
	defer vertex.Quit()

	window := w.NewWindow(800, 600, "Vertex Engine")

	for window.Loop() {
		draw.SetBackgroundColor([3]float32{0.2, 0.2, 0.2})
		draw.DrawLine(100, 100, 500, 200, [3]float32{1, 0, 0})
		draw.DrawRect(200, 200, 100, 100, [3]float32{0, 1, 1})

		if event.IsKeyPressedOnce(window, event.KeyA) {
			fmt.Println("A is pressed")
		}
	}
}

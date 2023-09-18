package draw

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
)

var drawn bool = false

func DrawLine(x1, y1, x2, y2 float32, color [3]float32) {
	if !drawn {
		gl.Begin(gl.LINES)
		gl.Color3f(color[0], color[1], color[2])
		gl.Vertex2f(x1, y1)
		gl.Vertex2f(x2, y2)
		gl.End()

		fmt.Println("[DEBUG] Line with coordinates", x1, y1, "to", x2, y2, "drawn")

		drawn = true
	}
}

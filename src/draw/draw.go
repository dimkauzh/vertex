package draw

import (
	"github.com/go-gl/gl/v2.1/gl"
)

func DrawLine(x1, y1, x2, y2 float32, color [3]float32) {
	gl.LineWidth(1)
	gl.Begin(gl.LINES)
	defer gl.End()

	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
}

func SetBackgroundColor(color [3]float32) {
	gl.ClearColor(color[0], color[1], color[2], 1)
}

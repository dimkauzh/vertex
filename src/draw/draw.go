package draw

import (
	"github.com/go-gl/gl/v2.1/gl"
)

func DrawLine(x1, y1, x2, y2 float32, color [3]float32) {
	gl.Begin(gl.LINES)
	defer gl.End() // Ensure glEnd() is called even if an error occurs.

	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x1, y1)
	gl.Vertex2f(x2, y2)
}

func DrawRect(x, y, height, width float32, color [3]float32) {
	gl.Begin(gl.QUADS)
	defer gl.End() // Ensure glEnd() is called even if an error occurs.

	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+width, y)
	gl.Vertex2f(x+width, y+height)
	gl.Vertex2f(x, y+height)
}

func SetBackgroundColor(color [3]float32) {
	gl.ClearColor(color[0], color[1], color[2], 1)
}

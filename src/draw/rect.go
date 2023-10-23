package draw

import "github.com/go-gl/gl/v2.1/gl"

func DrawRect(x, y, height, width float32, color [3]float32) {
	gl.Begin(gl.QUADS)
	defer gl.End() // Ensure glEnd() is called even if an error occurs.

	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2f(x, y)
	gl.Vertex2f(x+width, y)
	gl.Vertex2f(x+width, y+height)
	gl.Vertex2f(x, y+height)
}

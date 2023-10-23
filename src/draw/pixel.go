package draw

import "github.com/go-gl/gl/v2.1/gl"

func SetPixel(x, y int32, color [3]float32) {
	gl.Begin(gl.POINTS)
	defer gl.End()

	gl.Color3f(color[0], color[1], color[2])
	gl.Vertex2i(x, y)
}

package shape

import "github.com/dimkauzh/vertex/src/draw"

type Rect struct {
	width  float32
	height float32
	x      float32
	y      float32
	color  [3]float32
}

func (r *Rect) Draw() {
	draw.DrawRect(r.x, r.y, r.width, r.height, r.color)
}

func NewRect(x, y, width, height float32, color [3]float32) Rect {
	return Rect{width, height, x, y, color}
}

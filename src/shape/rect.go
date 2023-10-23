package shape

import (
	"github.com/dimkauzh/vertex/src/draw"
	"github.com/dimkauzh/vertex/src/vector"
)

type Rect struct {
	Pos    vector.Vector2D
	Width  float32
	Height float32
	Color  [3]float32
}

func (r *Rect) Draw() {
	draw.DrawRect(r.Pos.X, r.Pos.Y, r.Width, r.Height, r.Color)
}

func NewRect(x, y, width, height float32, color [3]float32) *Rect {
	return &Rect{vector.NewVector2D(x, y), width, height, color}
}

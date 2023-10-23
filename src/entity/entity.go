package entity

import (
    "github.com/dimkauzh/vertex/src/shape"
    "github.com/dimkauzh/vertex/src/vector"
)

type Entity struct {
    Pos    vector.Vector2D
    Width  float32
    Height float32
    Rect   *shape.Rect
}

func NewEntity(x, y, width, height float32) Entity {
    return Entity{vector.NewVector2D(x, y), width, height, nil}
}

func (e *Entity) NewRect(x, y, width, height float32, color [3]float32) {
    e.Rect = shape.NewRect(x, y, width, height, color)
}

// Deprecated: Won't be supported and will be deleted soon
func (e *Entity) SetRectColor(color [3]float32) {
    e.Rect.Color = color
}

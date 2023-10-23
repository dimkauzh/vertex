package vector

type Vector2D struct {
    X float32
    Y float32
}

func NewVector2D(x, y float32) Vector2D {
    return Vector2D{x, y}
}

package vector

type Vector3D struct {
    X float32
    Y float32
    Z float32
}

func NewVector3D(x, y, z float32) Vector3D {
    return Vector3D{x, y, z}
}

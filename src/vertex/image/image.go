package image

import (
	"image"
	"image/draw"
	_ "image/png"
	"log"
	"os"

	"github.com/go-gl/gl/v2.1/gl"
)

type Image struct {
	width   float32
	height  float32
	x       float32
	y       float32
	texture uint32
}

func LoadImage(file string, width, height, x, y float32) Image {
	imgFile, err := os.Open(file)
	if err != nil {
		log.Fatalf("texture %q not found on disk: %v\n", file, err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic(err)
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	return Image{width, height, x, y, texture}
}

func DrawImage(texture Image) {
	gl.BindTexture(gl.TEXTURE_2D, texture.texture)

	//gl.Color4f(1, 1, 1, 1)

	gl.Begin(gl.QUADS)

	gl.Normal3f(0, 0, 1)
	gl.TexCoord2f(0, 1) // Flip the Y-coordinate
	gl.Vertex3f(texture.x-texture.width, texture.y-texture.height, 0)
	gl.TexCoord2f(1, 1) // Flip the Y-coordinate
	gl.Vertex3f(texture.x+texture.width, texture.y-texture.height, 0)
	gl.TexCoord2f(1, 0) // Flip the Y-coordinate
	gl.Vertex3f(texture.x+texture.width, texture.y+texture.height, 0)
	gl.TexCoord2f(0, 0) // Flip the Y-coordinate
	gl.Vertex3f(texture.x-texture.width, texture.y+texture.height, 0)

	gl.End()
}

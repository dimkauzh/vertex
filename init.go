package vertex

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func Init() {
	err := glfw.Init()
	if err != nil {
		print("Error initialize")
		panic(err)
	}
}

func Quit() {
	glfw.Terminate()
}

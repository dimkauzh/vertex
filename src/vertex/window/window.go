package window

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

func NewWindow(width int, height int, title string) *glfw.Window {
	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		print("Error create window")
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}

func quit() {
	glfw.Terminate()
}

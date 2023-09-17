package window

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewWindow(width int, height int, title string) *glfw.Window {
	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		print("Error create window")
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}

func Loop(window *glfw.Window) bool {
	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
		return true
	}
	return false
}

package window

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewWindow(width int, height int, title string) *glfw.Window {
	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		fmt.Println("Error create window")
		panic(err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		fmt.Println("Error init gl")

		panic(err)

	}

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

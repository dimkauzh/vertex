package window

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Window struct {
	window *glfw.Window
	width  int
	height int
	title  string
}

func (w *Window) Loop() bool {
	for !w.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		return true
	}
	return false
}

func (w *Window) Refresh() {
	w.window.SwapBuffers()
	glfw.PollEvents()
}

func NewWindow(width int, height int, title string) Window {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

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
	gl.Viewport(0, 0, int32(width), int32(height))

	return Window{window, width, height, title}
}

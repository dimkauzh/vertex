package window

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const True int = 1
const False int = 0

var started bool = false

type Window struct {
	Window *glfw.Window
	Width  int
	Height int
	Title  string
}

func (w *Window) Loop() bool {
	for !w.Window.ShouldClose() {

		if started {
			w.Window.SwapBuffers()
		}

		glfw.PollEvents()
		gl.Clear(gl.COLOR_BUFFER_BIT)

		if !started {
			started = true
		}
		return true
	}
	return false
}

// Deprecated: No longer needed because of the integration into the Loop() method
func (w *Window) Refresh() {
	w.Window.SwapBuffers()
	glfw.PollEvents()
}

func NewCustomWindow(width, height int, title string, resizable int) Window {
	runtime.LockOSThread()

	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	if resizable == 1 {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else if resizable == 0 {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	} else {
		log.Fatalln("Wrong definition of resizable")
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)

	if err != nil {
		log.Fatalln("Error create window:", err)
	}

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatalln("Error init gl:", err)

	}
	gl.Ortho(0, float64(width), float64(height), 0, -1, 1)

	return Window{window, width, height, title}
}

func NewWindow(width, height int, title string) Window {
	return NewCustomWindow(width, height, title, False)
}

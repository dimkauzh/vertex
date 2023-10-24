package event

import (
	"github.com/dimkauzh/vertex/src/window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var keyPressed bool = false

func IsKeyPressed(window window.Window, key glfw.Key) bool {
	return window.Window.GetKey(key) == glfw.Press
}

func IsKeyPressedOnce(window window.Window, key glfw.Key) bool {
	if window.Window.GetKey(key) == glfw.Press {
		if !keyPressed {
			keyPressed = true
			return true
		}
	} else {
		keyPressed = false
	}
	return false
}

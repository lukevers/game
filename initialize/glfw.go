package initialize

import (
	"github.com/lukevers/go-gl/glfw/v3.1/glfw"
)

func Glfw() (w *glfw.Window) {
	// Initialize glfw
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	// Create glfw window
	w, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	return
}

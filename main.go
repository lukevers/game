package main

import (
	"github.com/apsdehal/go-logger"
	"github.com/lukevers/game/initialize"
	"github.com/lukevers/go-gl/glfw/v3.1/glfw"
	"github.com/lukevers/go-gl/text"
	"github.com/lukevers/go-gl/gl"
	"strconv"
)

// Logger
var log *logger.Logger

// Glfw Window
var window *glfw.Window

// Fonts
var fonts map[string]*text.Font

func main() {
	// Initialize logger
	log = initialize.Logger()

	// Initialize glfw and gl
	window = initialize.Glfw()
	initialize.Gl()

	// Defer termination to end
	defer glfw.Terminate()

	// Initialize fonts
	fonts = initialize.Fonts()

	// Run main loop
	loop()
}

func loop() {
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)


		// draw time for test
		drawDebug()


		// Poll for events
		glfw.PollEvents()

		// Swap Buffers
		window.SwapBuffers()
	}
}

func drawDebug() {
	time := strconv.FormatFloat(glfw.GetTime(), 'f', 0, 64) + " seconds"
	gl.Color4f(255, 255, 255, 1)
	fonts["AnonymiceBold"].Printf(5, 5, time)
}

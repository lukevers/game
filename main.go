package main

import (
	"github.com/apsdehal/go-logger"
	"github.com/lukevers/game/initialize"
	"github.com/lukevers/go-gl/glfw/v3.1/glfw"
	"github.com/lukevers/go-gl/text"
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
		// TODO

		// Poll for events
		glfw.PollEvents()

		// Swap Buffers
		window.SwapBuffers()
	}
}

package initialize

import (
	"github.com/lukevers/go-gl/gl"
	"github.com/lukevers/go-gl/glu"
)

func Gl() {
	// Initialize gl
	errno := gl.Init()
	if errno != gl.NO_ERROR {
		strerr, err := glu.ErrorString(errno)
		if err != nil {
			panic(err)
		} else {
			panic(strerr)
		}
	}

	// Set default background color to black
	gl.ClearColor(0, 0, 0, 1)
}

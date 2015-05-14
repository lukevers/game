package initialize

import (
	"github.com/lukevers/go-gl/text"
	"github.com/lukevers/game/fonts"
	"os"
)

func Fonts() (allfonts map[string]*text.Font) {
	// Loop over each font that we have and load it
	for name, file := range fonts.Fonts {
		font, err := loadFont(file, 14)
		if err == nil {
			allfonts[name] = font
		}
	}

	return
}

func loadFont(path string, scale int32) (*text.Font, error) {
	file, err := os.Open(fonts.Path + path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return text.LoadTruetype(file, scale, 32, 127, text.LeftToRight)
}

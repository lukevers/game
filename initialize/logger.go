package initialize

import (
	"github.com/apsdehal/go-logger"
	"os"
)

func Logger() *logger.Logger {
	// Create new logger
	log, err := logger.New("Logger", 1, os.Stdout)
	if err != nil {
		panic(err)
	}

	return log
}

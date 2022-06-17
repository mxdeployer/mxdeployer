package logfactory

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// TODO: Switch to lumberjack for file logging https://github.com/natefinch/lumberjack

func NewLogWriter() (io.Writer, error) {

	var baselogpath string

	if runtime.GOOS == "windows" {
		baselogpath = os.Getenv("ProgramData")
	} else {
		baselogpath = os.ExpandEnv("$HOME/.local/share")
	}

	logpath := filepath.Join(baselogpath, "/mxdeployer/logs/mxdeployer.log")
	logdir := filepath.Dir(logpath)

	if _, err := os.Stat(logdir); os.IsNotExist(err) {
		err := os.MkdirAll(logdir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	logFile, err := os.Create(logpath)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Logs written to: %s\n", logpath)

	return io.MultiWriter(logFile, os.Stdout), nil
}

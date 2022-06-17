package logfactory

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/natefinch/lumberjack.v2"
)

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

	rollinglog := lumberjack.Logger{
		Filename:  logpath,
		MaxAge:    31,
		LocalTime: true,
	}

	fmt.Printf("Logs written to: %s\n", logpath)

	return io.MultiWriter(&rollinglog, os.Stdout), nil
}

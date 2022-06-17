package logfactory

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// TODO: Make new writer type that support io.WriteCloser, close logFile

func NewLogWriter() (io.Writer, error) {

	home, _ := os.UserHomeDir()
	logpath := filepath.Join(home, ".local/share/mxdeployer/logs/mxdeployer.log")
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

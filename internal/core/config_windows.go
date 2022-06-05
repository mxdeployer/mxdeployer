//go:build windows

package core

import "os"

func ConfigPath() string {
	return os.ExpandEnv("$PROGRAMDATA\\mxdeployer\\config.json")
}

func ApplicationBasePath() string {
	return os.ExpandEnv("$PROGRAMFILES")
}

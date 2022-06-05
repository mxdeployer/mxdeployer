//go:build !windows

package core

import "os"

func ConfigPath() string {

	if os.Geteuid() == 0 {
		return "/usr/local/share/mxdeployer/config.json"
	}

	return os.ExpandEnv("$HOME/.local/share/mxdeployer/config.json")
}

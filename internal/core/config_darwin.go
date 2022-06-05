//go:build darwin

package core

import "os"

func ApplicationBasePath() string {
	if os.Geteuid() == 0 {
		return "/Applications"
	}

	return os.ExpandEnv("$HOME/Applications")
}

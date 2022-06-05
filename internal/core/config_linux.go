//go:build linux

package core

func ApplicationBasePath() string {
	if os.Geteuid() == 0 {
		return "/usr/local"
	}

	return os.ExpandEnv("$HOME/.local")
}

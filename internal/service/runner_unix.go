//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package service

import (
	"os"
	"os/signal"
	"syscall"
)

func Run(s Service) {
	s.Start()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		done <- true
	}()

	<-done

	s.Stop()
}

//go:build windows

package service

import (
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)

type envelope struct {
	serviceImpl Service
}

func Run(s Service) {

	isService, _ := svc.IsWindowsService()

	if isService {
		svc.Run(s.Name(), &envelope{serviceImpl: s})
	} else {
		debug.Run(s.Name(), &envelope{serviceImpl: s})
	}
}

func (se *envelope) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	se.serviceImpl.Start()
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	running := true
	for running {
		change := <-r
		switch change.Cmd {
		case svc.Interrogate:
			changes <- change.CurrentStatus
			// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
			time.Sleep(100 * time.Millisecond)
			changes <- change.CurrentStatus
		case svc.Stop, svc.Shutdown:
			se.serviceImpl.Stop()
			running = false
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

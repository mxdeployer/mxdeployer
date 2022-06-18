package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/mxdeployer/mxdeployer/internal/core"
	"github.com/mxdeployer/mxdeployer/internal/logfactory"
	"github.com/mxdeployer/mxdeployer/internal/models"
	"github.com/mxdeployer/mxdeployer/internal/service"
)

func main() {
	logwriter, err := logfactory.NewLogWriter()

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(logwriter)

	service.Run(&mxdeployservice{})
}

type mxdeployservice struct {
	queue  core.NotificationQueue
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func (s *mxdeployservice) Name() string {
	return "mxdeployservice"
}

func (s *mxdeployservice) Start() {

	log.Printf("Starting... %d\n", os.Getpid())

	cfg := core.LoadConfig()

	s.queue = *core.NewNotificationQueue(cfg.AzServiceBusConStr, cfg.Host)

	var cancelCtx context.Context
	cancelCtx, s.cancel = context.WithCancel(context.Background())

	go s.queue.Process(cancelCtx, s.NotificationReceived)
	log.Println("Started.")
}

func (s *mxdeployservice) Stop() {
	log.Println("Stopping...")
	s.cancel()
	log.Println("Waiting for deployments...")
	s.wg.Wait()
	log.Println("Stopped.")
}

func (s *mxdeployservice) NotificationReceived(dn models.DeploymentNotification) {

	s.wg.Add(1)
	go func(dn models.DeploymentNotification) {
		log.Printf("We got one - %s %s\n", dn.AppName, dn.Ref)
		defer s.wg.Done()

		// TODO: Download artifact from dn.Url
		// TODO: Run stop script (if present)
		// TODO: Extract artifact to destination
		// TODO: Deploy system environment variables
		// TODO: Deploy app secrets
		// TODO: Run install script (if present)
		// TODO: Run start script (if present)

	}(dn)

}

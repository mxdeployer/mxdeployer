package main

import (
	"context"
	"log"

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
}

func (s *mxdeployservice) Name() string {
	return "mxdeployservice"
}

func (s *mxdeployservice) Start() {

	log.Println("Starting...")

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
	log.Println("Stopped.")
}

func (s *mxdeployservice) NotificationReceived(dn models.DeploymentNotification) {

	log.Printf("We got one - %s %s\n", dn.AppName, dn.Ref)
}

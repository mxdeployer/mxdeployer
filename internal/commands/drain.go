package commands

import (
	"fmt"

	"github.com/mxdeployer/mxdeployer/internal/core"
)

type Drain struct {
	constr string
	host   string
}

func NewDrain(argQueue *core.StringQueue) core.Command {
	cfg := core.LoadConfig()
	return &Drain{cfg.AzServiceBusConStr, argQueue.DequeueOrDefault(cfg.Host)}
}

func (cmd *Drain) Run() error {

	queue := core.NewNotificationQueue(cmd.constr, cmd.host)

	for msg, _ := queue.Receive(5); msg != nil; msg, _ = queue.Receive(5) {
		fmt.Print(".")
	}

	fmt.Println()

	return nil
}

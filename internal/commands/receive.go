package commands

import (
	"encoding/json"
	"fmt"

	"github.com/mxdeployer/mxdeployer/internal/core"
)

type Receive struct {
	constr string
	host   string
}

func NewReceive(argQueue *core.StringQueue) core.Command {
	cfg := core.LoadConfig()
	return &Receive{cfg.AzServiceBusConStr, argQueue.DequeueOrDefault(cfg.Host)}
}

func (cmd *Receive) Run() error {

	queue := core.NewNotificationQueue(cmd.constr, cmd.host)

	not, err := queue.Receive(10)

	if err != nil {
		return nil
	}

	if not == nil {
		fmt.Println("No messages waiting!")
		return nil
	}

	notJson, err := json.MarshalIndent(not, "", "    ")

	if err != nil {
		return err
	}

	fmt.Println(string(notJson))

	return nil
}

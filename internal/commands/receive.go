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

	dn, err := queue.Receive(10)

	if err != nil {
		return nil
	}

	if dn == nil {
		fmt.Println("No messages waiting!")
		return nil
	}

	dnJson, err := json.MarshalIndent(dn, "", "    ")

	if err != nil {
		return err
	}

	fmt.Println(string(dnJson))

	return nil
}

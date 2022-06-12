package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mxdeployer/mxdeployer/internal/core"
	"github.com/mxdeployer/mxdeployer/internal/models"
)

type Send struct {
	jsonPath string
}

func NewSend(argQueue *core.StringQueue) core.Command {
	return &Send{argQueue.Dequeue()}
}

func (cmd *Send) Run() error {

	config := core.LoadConfig()

	queue := core.NewNotificationQueue(config.AzServiceBusConStr, config.Host)

	content, err := ioutil.ReadFile(cmd.jsonPath)

	if err != nil {
		return err
	}

	var notification models.DeploymentNotification
	err = json.Unmarshal(content, &notification)

	if err != nil {
		return err
	}

	err = queue.Send(notification)

	if err != nil {
		return err
	}

	fmt.Println("OK.")

	return nil
}

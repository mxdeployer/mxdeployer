package commands

import (
	"encoding/json"
	"fmt"

	"github.com/mxdeployer/mxdeployer/internal/core"
	"github.com/mxdeployer/mxdeployer/internal/models"
)

func NewInit() core.Command {
	return &Init{}
}

type Init struct {
}

func (cmd *Init) Run() error {

	not := models.NewDeploymentNotification()
	notJson, err := json.MarshalIndent(not, "", "    ")

	if err != nil {
		return err
	}

	fmt.Println(string(notJson))

	return nil
}

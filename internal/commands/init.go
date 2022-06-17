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

	dn := models.NewDeploymentNotification()
	dnJson, err := json.MarshalIndent(dn, "", "    ")

	if err != nil {
		return err
	}

	fmt.Println(string(dnJson))

	return nil
}

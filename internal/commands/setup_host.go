package commands

import (
	"fmt"

	"github.com/mxdeployer/mxdeployer/internal/core"
)

type SetupHost struct {
	config core.Config
}

func NewSetupHost(org string, ascs string, asbcs string, host string) core.Command {
	return &SetupHost{*core.NewConfig(org, ascs, asbcs, host)}
}

func (cmd *SetupHost) Run() error {
	err := core.SaveConfig(cmd.config)

	if err != nil {
		return err
	}

	fmt.Printf("Configuration saved to: %s\n", core.ConfigPath())

	return nil
}

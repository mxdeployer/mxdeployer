package core

import (
	"log"
	"os"
)

func NewSetupCommand() Command {
	return &SetupCommand{}
}

type SetupCommand struct {
}

func (cmd *SetupCommand) Run() {

	azstorageconstr := os.Getenv("MXDEPLOYER_AZSTORAGE")
	azservicebusconstr := os.Getenv("MXDEPLOYER_AZSERVICEBUS")

	log.Print(azstorageconstr)
	log.Print(azservicebusconstr)

	log.Fatal("Not implemented yet!")
}

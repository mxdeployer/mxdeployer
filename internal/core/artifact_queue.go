package core

import (
	"fmt"

	"github.com/mxdeployer/mxdeployer/internal/models"
)

type ArtifactQueue struct {
	constr string
}

func NewArtifactQueue(constr string) *ArtifactQueue {
	return &ArtifactQueue{constr: constr}
}

func (queue ArtifactQueue) Send(artifact *models.Artifact) {
	fmt.Printf("Sending '%s' ...\n", artifact.Url)
}

func (queue ArtifactQueue) Receive() {
	fmt.Println("Receiving...")
}

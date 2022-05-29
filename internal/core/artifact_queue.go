package core

import "fmt"

type ArtifactQueue struct {
	constr string
}

func NewArtifactQueue(constr string) *ArtifactQueue {
	return &ArtifactQueue{constr: constr}
}

func (queue ArtifactQueue) Send() {
	fmt.Println("Sending...")
}

func (queue ArtifactQueue) Receive() {
	fmt.Println("Receiving...")
}

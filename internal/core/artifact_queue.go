package core

type ArtifactQueue struct {
	constr string
}

func NewArtifactQueue(constr string) *ArtifactQueue {
	return &ArtifactQueue{constr: constr}
}

package models

type Artifact struct {
	url string
}

func NewArtifact(url string) *Artifact {
	return &Artifact{url: url}
}

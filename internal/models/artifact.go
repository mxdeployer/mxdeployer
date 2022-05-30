package models

type Artifact struct {
	Url string
}

func NewArtifact(url string) *Artifact {
	return &Artifact{Url: url}
}

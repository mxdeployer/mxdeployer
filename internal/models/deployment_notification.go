package models

type DeploymentNotification struct {
	Url         string            `json:"url"`
	Ref         string            `json:"ref"`
	Host        string            `json:"host"`
	AppName     string            `json:"appName"`
	Environment map[string]string `json:"environment"`
	AppSecrets  map[string]string `json:"appSecrets"`
}

func NewDeploymentNotification() *DeploymentNotification {
	return &DeploymentNotification{
		Url:         "",
		Ref:         "",
		Host:        "",
		AppName:     "",
		Environment: make(map[string]string),
		AppSecrets:  make(map[string]string),
	}
}

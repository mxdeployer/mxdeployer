package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Org                string `json:"org"`
	AzStorageConStr    string `json:"az-storage-con-str"`
	AzServiceBusConStr string `json:"az-service-bus-con-str"`
	Host               string `json:"host"`
	AppBasePath        string `json:"app-base-path"`
}

func NewConfig(org string, ascs string, asbcs string, host string) *Config {
	return &Config{org, ascs, asbcs, host, ApplicationBasePath()}
}

func LoadConfig() Config {

	content, err := ioutil.ReadFile(ConfigPath())

	if err != nil {
		return *NewConfig("", "", "", "")
	} else {
		var config Config
		json.Unmarshal(content, &config)
		return config
	}
}

func SaveConfig(config Config) error {

	configpath := ConfigPath()
	dir, _ := filepath.Split(configpath)

	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(config, "", "    ")

	if err != nil {
		return err
	}

	f, err := os.Create(configpath)

	if err != nil {
		return err
	}

	defer f.Close()

	f.Write(bytes)

	return nil
}

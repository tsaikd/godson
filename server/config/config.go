package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/tsaikd/godson/server/api/bitbucket"
)

type Config struct {
	Bitbuckets bitbucket.Configs `json:"bitbuckets,omitempty"`
}

func NewConfigFromFile(filename string) (retconf *Config, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	return NewConfigFromData(data)
}

func NewConfigFromData(data []byte) (retconf *Config, err error) {
	config := Config{}
	if err = json.Unmarshal(data, &config); err != nil {
		return
	}

	if err = InitConfig(&config); err != nil {
		return
	}

	retconf = &config
	return
}

func InitConfig(config *Config) (err error) {
	for i, _ := range config.Bitbuckets {
		if err = bitbucket.InitConfig(&config.Bitbuckets[i]); err != nil {
			return
		}
	}

	return
}

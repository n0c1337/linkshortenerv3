package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Address string
}

func ParseConfigFromJSONFile(fileName string) (c *Config, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	c = new(Config)
	json.NewDecoder(f).Decode(c)

	return
}

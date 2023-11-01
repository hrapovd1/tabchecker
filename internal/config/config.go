/*
package config includes types and methods for app config
*/
package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
)

var ErrNoConfig = errors.New("Config not provided")

type Config struct {
	Left  string `json:"left"`
	Right string `json:"right"`
	Table Table  `json:"table,omitempty"`
}

type Table struct {
	Name string `json:"name"`
}

func NewConfig() (Config, error) {
	config := Config{}
	configPath, err := config.getConfig()
	if err != nil {
		return config, err
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	configRaw, err := io.ReadAll(configFile)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(configRaw, &config); err != nil {
		return config, err
	}
	return config, nil
}

// return config path from args or error
func (c *Config) getConfig() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", ErrNoConfig
	}
	return filepath.Abs(args[1])
}

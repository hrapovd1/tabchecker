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

// Config struct for app configuration, must be in json:
//
//		{
//			"left": {
//	            "type": "type",
//	            "dsn": "dsn"
//	        },
//			"right": {
//	            "type": "type",
//	            "dsn": "dsn"
//	        },
//			"table": {
//				"name": "name",
//				"fields": {
//					"colonName1": "type",
//					"colonName2": "type"
//				}
//			}
//		}
type dsnConfig struct {
	Type string `json:"type"`
	DSN  string `json:"dsn"`
}
type Config struct {
	Left  dsnConfig `json:"left"`
	Right dsnConfig `json:"right"`
	Table Table     `json:"table,omitempty"`
}

type Table struct {
	Name string `json:"name"`
}

// NewConfig read cmd args and wait config path
// from which it generates app config
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
	defer configFile.Close()

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

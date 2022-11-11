package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	l "github.com/zanshin/atempo/pkg/logger"
)

var (
	configFilePath string
)

// DbConfig contains the database settings.
type DbConfig struct {
	Host       string `json:"host"`
	PortNumber int    `json:"port"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
	DBName     string `json:"name"`
}

// Config contains the main configuration settings for the application.
type Config struct {
	BatchInsertSeconds int      `json:"batchInsertSeconds"`
	PortNumber         int      `json:"port"`
	DbConfig           DbConfig `json:"database"`
}

func ReadConfig(configFilePath string) Config {
	config := Config{}
	configFile, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		l.Error.Fatal("Unable to read config file: ", err)
	}

	if err = json.Unmarshal(configFile, &config); err != nil {
		l.Error.Fatal("Unable to unmarshal configFile into config: ", err)
	}

	if err = validateConfig(config); err != nil {
		l.Error.Fatal("Unable to validate config: ", err)
	}

	return config
}

func validateConfig(config Config) error {
	if config.BatchInsertSeconds < 1 {
		return fmt.Errorf(
			"BatchInsertSeconds cannot be less than 1, %d was given.",
			config.BatchInsertSeconds,
		)
	}
	if config.PortNumber < 0 || config.PortNumber > 65535 {
		return fmt.Errorf(
			"PortNumber must be between 0 and 65535, %d was given.",
			config.PortNumber,
		)
	}
	if config.DbConfig.PortNumber < 0 || config.DbConfig.PortNumber > 65535 {
		return fmt.Errorf(
			"The database port must be between 0 and 65535, %d was given.",
			config.DbConfig.PortNumber,
		)
	}
	return nil
}

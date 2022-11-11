package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zanshin/atempo/pkg/dbconf"
	l "github.com/zanshin/atempo/pkg/logger"
)

var (
	configFilePath string
)

// Config contains the main configuration settings for the application.
type Config struct {
	BatchInsertSeconds int             `json:"batchInsertSeconds"`
	Port               int             `json:"port"`
	DbConfig           dbconf.DbConfig `json:"database"`
}

// Flags, config file path, setup logging
func init() {
	goPath := os.Getenv("GOPATH")
	defaultConfigPath := fmt.Sprintf("%s/src/github.com/zanshin/atempo/config.json", goPath)
	flag.StringVar(&configFilePath, "config", defaultConfigPath, "path to config.json")

	l.Setup("wa.log")
}

func main() Config {
	// Read the config, initialize the database and listen for records.
	flag.Parse()
	return readConfig(configFilePath)
}

func readConfig(configFilePath string) Config {
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
	if config.Port < 0 || config.Port > 65535 {
		return fmt.Errorf(
			"Port must be between 0 and 65535, %d was given.",
			config.Port,
		)
	}
	if config.DbConfig.Port < 0 || config.DbConfig.Port > 65535 {
		return fmt.Errorf(
			"The database port must be between 0 and 65535, %d was given.",
			config.DbConfig.Port,
		)
	}
	return nil
}

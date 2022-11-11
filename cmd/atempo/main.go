package main

import (
	"flag"
	"fmt"

	"github.com/zanshin/atempo/pkg/config"
	l "github.com/zanshin/atempo/pkg/logger"
)

var (
	configFilePath string
)

// Flags, config file path, setup logging
func init() {
	// goPath := os.Getenv("GOPATH")
	defaultConfigPath := fmt.Sprintf("/Users/mark/code/go/src/github.com/zanshin/atempo/config.json")
	flag.StringVar(&configFilePath, "config", defaultConfigPath, "path to config.json")

	l.Setup("atempo.log")
}

func main() {
	// Read the config, initialize the database and listen for records.
	flag.Parse()
	conf := config.ReadConfig(configFilePath)

	l.Info.Println("Configuration read from file system.")
	fmt.Println(conf)
}

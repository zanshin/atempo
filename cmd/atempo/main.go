package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zanshin/atempo/internal/app"
	"github.com/zanshin/atempo/internal/config"
	l "github.com/zanshin/atempo/internal/logger"
)

var (
	configFilePath string
)

func main() {
	goPath := os.Getenv("GOPATH")
	defaultConfigPath := fmt.Sprintf("%s/src/github.com/zanshin/atempo/config.json", goPath)

	setup := flag.Bool("s", false, "Perform inital app setup, including database")
	listen := flag.Bool("l", false, "Listen for visitor events")
	flag.StringVar(&configFilePath, "config", defaultConfigPath, "path to config.json")

	flag.Parse()
	conf := config.ReadConfig(configFilePath)
	l.Info.Println("Configuration read from file system.")

	l.Setup("atempo.log")

	if *setup {
		app.RunSetup(conf)
	}

	if *listen {
		app.Run(conf)
	}
}

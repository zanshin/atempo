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
	l.Setup("logs/atempo.log")

	l.Info.Println("Starting Atempo")
	goPath := os.Getenv("GOPATH")
	defaultConfigPath := fmt.Sprintf("%s/src/github.com/zanshin/atempo/config.json", goPath)

	setup := flag.Bool("s", false, "Perform inital app setup, including database")
	listen := flag.Bool("l", false, "Listen for visitor events")
	path := flag.String("p", defaultConfigPath, "Path to configuration file")

	flag.Parse()

	l.Info.Printf("Configuration path %q", *path)
	conf := config.ReadConfig(*path)
	l.Info.Println("Configuration read from file system.")

	if *setup {
		app.RunSetup(conf)
	}

	if *listen {
		app.Run(conf)
	}
}

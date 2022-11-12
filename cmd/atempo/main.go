package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/zanshin/atempo/pkg/config"
	l "github.com/zanshin/atempo/pkg/logger"
	"github.com/zanshin/atempo/pkg/model"
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

	// Connect to the database
	db, err := model.DBConnect("", conf.DbConfig)
	if err != nil {
		l.Error.Fatal("Connection to MySQL failed. Exiting.")
	}

	l.Info.Println("Successfully connected to MySQL")

	// Create database if it doesn't already exist
	dbname := "atempo"
	res, err := model.DBCreate(db, dbname)
	if err != nil {
		l.Error.Fatal("Database creation failed. Exiting")
	}

	// Test result of creation for number of row affected.
	no, err := res.RowsAffected()
	if err != nil {
		l.Error.Printf("Error %s when fetching affected rows\n", err)
		l.Error.Fatal("Unable to get affected rows. Exiting")
	}

	l.Info.Printf("Rows affected: %d\n", no)

	// Set connection pool options
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	// Verify the connection
	err = model.DBPing(db)
	if err != nil {
		l.Error.Fatal("Unable to ping database. Exiting")
	}

	l.Info.Printf("Connected to DB %s successfully\n", dbname)
}

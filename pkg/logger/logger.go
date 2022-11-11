package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

const (
	Red    string = "\033[;31m"
	Yellow string = "\033[;33m"
	Green  string = "\033[;32m"
	Reset  string = "\033[;0m"
)

func Setup(name string) {

	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// es := fmt.Sprintf(Red, "Error   : ", Reset)
	es := fmt.Sprintf("%s%s%s", string(Red), "Error   : ", string(Reset))
	ws := fmt.Sprintf("%s%s%s", string(Yellow), "Warning : ", string(Reset))
	is := fmt.Sprintf("%s%s%s", string(Green), "Info    : ", string(Reset))

	Info = log.New(file, is, log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(file, ws, log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, es, log.Ldate|log.Ltime|log.Lshortfile)
}

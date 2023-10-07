package logger

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	DebugLogger   *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	year, month, day := time.Now().Date()

	file_name := fmt.Sprintf("%s %s %s", strconv.Itoa(day), month.String(), strconv.Itoa(year))
	file_name += ".log"
	file, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	WarningLogger = log.New(file, "WARNING : ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(file, "DEBUG : ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
}

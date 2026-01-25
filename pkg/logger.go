package pkg

import (
	"log"
	"os"
)

var InfoLogger *log.Logger
var ErrorLogger *log.Logger

func InitLogger() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

package logger

import (
	"fmt"
	"log"
	"os"
)

var filename string = "airgabe.log"

func Info(message []string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println(fmt.Sprintf("Request: %s", message))
}

func Error(message string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	log.Println(fmt.Sprintf("ERROR: %s", message))
}

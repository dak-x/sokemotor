package utils

import "log"

func HandleFatalError(logMessage string, err error) {
	if err != nil {
		log.Fatalf("%s %s", logMessage, err)
	}
}

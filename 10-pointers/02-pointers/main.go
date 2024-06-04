package main

import (
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	profaneWords := [3]string{"fubb", "shiz", "witch"}

	for _, profaneWord := range profaneWords {
		censoredWord := ""
		if profaneWord == "fubb" {
			censoredWord = "****"
		} else if profaneWord == "shiz" {
			censoredWord = "****"
		} else {
			censoredWord = "*****"
		}

		messageVal = strings.ReplaceAll(messageVal, profaneWord, censoredWord)
	}
	*message = messageVal
}

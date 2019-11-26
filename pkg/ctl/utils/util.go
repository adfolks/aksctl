package utils

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fatih/color"
)

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count a
// confirmations. If the input is not recognized, it will ask again. The function does not reurn
// until it gets a valid response from the user. Typically, you should use fmt to print out a quetion
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func AskForConfirmation(okayResponses []string, nokayResponses []string, message string) bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	if ContainsString(okayResponses, response) {
		return true
	} else if ContainsString(nokayResponses, response) {
		return false
	} else {
		color.Blue(message)
		return AskForConfirmation(okayResponses, nokayResponses, message)
	}
}

// You might want to put the following two functions in a separate utility package.

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}

// containsString returns true iff slice contains element
func ContainsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func stringToMap(data string) []map[string]interface{} {
	var value []map[string]interface{}

	err := json.Unmarshal([]byte(data), &value)

	if err != nil {
		panic(err)
	}
	return value
}

func FilterStringMap(data string, key string) []string {
	mapdata := stringToMap(data)
	var slice []string
	for _, mapItem := range mapdata {
		slice = append(slice, mapItem[key].(string))
	}
	return slice
}

package confirm

import (
	"fmt"
	"log"

	"github.com/adfolks/aksctl/pkg/ctl/utils"
	"github.com/fatih/color"
)

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func AskForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if utils.ContainsString(okayResponses, response) {
		return true
	} else if utils.ContainsString(nokayResponses, response) {
		return false
	} else {
		color.Blue("Please type yes or no and then press enter:")
		return AskForConfirmation()
	}
}



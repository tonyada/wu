package wu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// TimeoutYesOrNo prompts the user for a yes/no answer with a timeout.
// Returns the defaultResult if the timeout expires or an invalid input is given.
func TimeoutYesOrNo(desc string, timeout time.Duration, defaultResult bool) bool {
	fmt.Printf("%s [y/n]: ", desc)

	// Use a buffered reader for more efficient input reading.
	reader := bufio.NewReader(os.Stdin)

	// Create a channel to receive the user's input.
	inputChan := make(chan string)

	// Launch a goroutine to read user input.
	go func() {
		input, err := reader.ReadString('\n')
		if err != nil {
			inputChan <- "" // Send empty string on error
			return
		}
		inputChan <- strings.TrimSpace(strings.ToLower(input))
	}()

	// Set a timer for the timeout.
	timer := time.After(timeout)

	// Wait for either the user input or the timeout.
	select {
	case input := <-inputChan:
		switch input {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			return defaultResult // Handle invalid input
		}
	case <-timer:
		fmt.Println("\nTimeout!")
		return defaultResult
	}
}

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

	done := make(chan bool, 1)
	var result bool

	go func() {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err == nil {
			input = strings.TrimSpace(strings.ToLower(input))
			switch input {
			case "y", "yes":
				result = true
			case "n", "no":
				result = false
			default:
				result = defaultResult
			}
		} else {
			result = defaultResult
		}
		done <- true
	}()

	select {
	case <-done:
		return result
	case <-time.After(timeout):
		fmt.Println("\nTimeout!")
		return defaultResult
	}
}

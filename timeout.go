package wu

import (
	"fmt"
	"time"
)

// Yes or no with TimeOut TimeoutYesOrNo("select", 6, false)
func TimeoutYesOrNo(desc string, timeoutSec int, defaultResult bool) bool {
	print(desc, " [y/n]")
	result := defaultResult
	// scan user input
	userInputEndSignal := make(chan bool, 1)
	go func() {
		userInput := ""
		fmt.Scanf("%s", &userInput)
		// println("userInput: ", userInput)
		if userInput == "y" { //@ must be 'y' to YES
			result = true
		} else if userInput == "n" {
			result = false
		}
		userInputEndSignal <- true
	}()
	// set timeout
	timeout := time.After(time.Duration(timeoutSec) * time.Second)
	isEnd := false
	for {
		select {
		case <-userInputEndSignal:
			isEnd = true // input running out
		case <-timeout:
			isEnd = true // timeout
		default:
			print(".")
		}
		if isEnd {
			break
		}
		time.Sleep(time.Second) // loop every sec
	}
	return result
}

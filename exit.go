package wu

import (
	"fmt"
	"os"
)

// Global common functions
// Exit with message | ExitWithMsg("sth is wrong")
func ExitWithMsg(msg string) { fmt.Println(msg); os.Exit(1) }

// Exit with desc | ExitWithDesc("sth is wrong")
func ExitWithDesc(desc string) { ExitWithMsg(desc) }

// Exit with message
func ErrExit(msg string) { fmt.Println(msg); os.Exit(1) }

// Exit with message and exit code
func ErrExit2(msg string, i int) { fmt.Println(msg); os.Exit(i) }

// os.Exit
func Exit(i int) { os.Exit(i) }

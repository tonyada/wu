package user

import (
	"os"
	"os/user"
	"runtime"
	. "wu"
	"wu/os/shell"
)

// Home returns the home directory for the executing user.
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	// darwin
	u, err := user.Current()
	if err == nil {
		return u.HomeDir, nil
	}
	// cross compile support
	if runtime.GOOS == "windows" {
		return homeWindows()
	}
	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}
	// If that fails, try the shell
	return shell.Run("echo -n $HOME")
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", ErrNew("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}
	return home, nil
}

// return current system username
func CurrentUsername() string {
	username, err := Username()
	if Err(err) {
		return "err"
	}
	return username
}

// return current system username
func Username() (string, error) {
	u, err := user.Current()
	if err == nil {
		return u.Username, nil
	}
	// cross compile support
	if runtime.GOOS == "windows" {
		return usernameWindows()
	}
	// Unix-like system, so just assume Unix
	return usernameUnix()
}

// TODO: not tested
func usernameWindows() (string, error) {
	u, err := user.Current()
	if err == nil {
		return u.Username, nil
	}
	return "", ErrNew("TODO: need to code")
}

func usernameUnix() (string, error) {
	// First prefer the HOME environmental variable
	if username := os.Getenv("USER"); username != "" {
		return username, nil
	}
	return shell.Run("echo -n $USER")
}

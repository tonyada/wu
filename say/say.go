package say

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"wu/str"
)

// command line say func osx only
// speak without wait
func Say(s ...string) error {
	if runtime.GOOS == "darwin" {
		return exec.Command("say", s...).Start()
	}
	if runtime.GOOS == "windows" {
		return errors.New("not supported")
	}
	return nil
}

func SayInteractive(s ...string) error {
	if runtime.GOOS == "darwin" {
		cli := str.Fields(fmt.Sprintf("-i %s", s))
		cmd := exec.Command("say", cli...)
		// Create stdout, stderr streams of type io.Reader
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()
		go func() {
			_, _ = io.Copy(os.Stdout, stdout)
		}()
		go func() {
			_, _ = io.Copy(os.Stderr, stderr)
		}()
		// cmd.Stdout = os.Stdout
		// cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	return nil
}

// waiting for saying finished
func SayWait(s ...string) error {
	if runtime.GOOS == "darwin" {
		return exec.Command("say", s...).Run()
	}
	if runtime.GOOS == "windows" {
		// c:\\Windows\\tts.exe -f 10 -v
		// c:\\Windows\\System32\\WindowsPowerShell\\v1.0\\PowerShell.exe -Command "Add-Type –AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('hello');"
		// mshta vbscript:Execute("CreateObject(""SAPI.SpVoice"").Speak(""你好"")(window.close)")
		// command := Sprintf(`/C c:\\Windows\\System32\\WindowsPowerShell\\v1.0\\PowerShell.exe -Command "Add-Type –AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('%v');"`, s)
		// command := `/C c:\\Windows\\System32\\WindowsPowerShell\\v1.0\\PowerShell.exe -Command "Add-Type –AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('hello');"`
		// command := "PowerShell -Command \"Add-Type –AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('你好');\""
		command := "`/C c:\\tts.exe -f 10 -v 0`"
		cli := strings.Split(command, " ")
		cli = append(cli, fmt.Sprintf("\"%v\"", s))
		// Println(cli)
		// app name = cli[0] cli = cli[1:]
		return exec.Command("c:\\Windows\\System32\\cmd.exe", cli...).Run()
	}
	return nil
}

// say with print out
func SayWithPrint(s ...string) error {
	fmt.Println(s)
	if runtime.GOOS == "darwin" {
		return exec.Command("say", s...).Run()
	}
	return nil
}

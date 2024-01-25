package shell

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

// type ShellCmd struct {
// 	*exec.Cmd
// 	Running bool
// 	Timeout time.Duration
// 	Pid     int
// 	Uid     int
// }

// // create a new instance
// func Command(cli string) *ShellCmd {
// 	args := strings.Split(cli, " ")
// 	exe_name := args[0]
// 	exe_args := args[1:]
// 	return &ShellCmd{exec.Command(exe_name, exe_args...), false, 0, 0, 0}
// }

// TODO: some cmd not working
// osascript -e 'tell app "Xcode" to quit'

func Command(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}

func FindExe(name string) (string, error) {
	return exec.LookPath(name)
}

func Cmd(command string) *exec.Cmd {
	// split command to exe and args... $ls -l /Users to []string{"ls", "-l", "/Users"}
	args := strings.Split(command, " ")
	// app name = args[0] args = args[1:]
	return exec.Command(args[0], args[1:]...)
}

// cmd := exec.Command("ls", "-l")
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(out.String())
func Exec(cli string) ([]byte, error) {
	return Cmd(cli).Output()
}

// run without wait
func Run(cli string) (string, error) {
	out, err := Cmd(cli).Output()
	return string(out), err
}

// run executes a program and waits for it to finish. The stdin, stdout, and
// stderr of noti are passed to the program.
func RunWait(cli string) error {
	args := strings.Split(cli, " ")
	// app name = args[0] args = args[1:]
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunDir(workDir string, cli string) (string, error) {
	cmd := Cmd(cli)
	// set working dir
	cmd.Dir = workDir
	// Logf("Run Path: %s Working dir: %s", cmd.Path, cmd.Dir)

	output, err := cmd.Output()
	return string(output), err
}

func Run2(cli string, isOutput bool) (string, error) {
	cmd := Cmd(cli)
	if isOutput {
		output, err := cmd.Output()
		return string(output), err
	} else {
		err := cmd.Start()
		return "", err
	}
	// cmd.Process.Release()
}

// start a cmd and return it as started
func CmdStart(cli string) (*exec.Cmd, error) {
	cmd := Cmd(cli)
	err := cmd.Start()
	return cmd, err
}

// stop given cmd
func CmdStop(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}

// kill given cmd
func Kill(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}

// A simpler version without select and channels.
// func RunTimeout() {
//     cmd := exec.Command("cat", "/dev/urandom")
//     cmd.Start()
//     timer := time.AfterFunc(1*time.Second, func() {
//         err := cmd.Process.Kill()
//         if err != nil {
//             panic(err) // panic as can't kill a process.
//         }
//     })
//     err := cmd.Wait()
//     timer.Stop()

//     // read error from here, you will notice the kill from the
//     fmt.Println(err)
// }

// even shorter version, and very straight forward.
// BUT, possibly having tons of hanging goroutines if timeout is long.
// func RunTimeout() {
//     cmd := exec.Command("cat", "/dev/urandom")
//     cmd.Start()
//     go func(){
//         time.Sleep(timeout)
//         cmd.Process.Kill()
//     }()
//     return cmd.Wait()
// }

func RunTimeout(cli string, sec_to_timeout time.Duration) (string, error) {
	cmd := Cmd(cli)
	if err := cmd.Start(); err != nil {
		return "", err
	}
	// wait or timeout
	cdone := make(chan error, 1)
	go func() {
		cdone <- cmd.Wait()
	}()
	select {
	case <-time.After(sec_to_timeout * time.Second):
		if err := cmd.Process.Kill(); err != nil {
			// failed to kill
			fmt.Printf("timed out %v secs, but %v is failed to kill. Err: %v\n", sec_to_timeout, cli, err)
		}
		<-cdone // allow goroutine to exit
		// process killed
		fmt.Printf("%v is killed because timed out %v secs\n", cli, sec_to_timeout)
		output, err := cmd.Output()
		return string(output), err
	case <-cdone:
		output, err := cmd.Output()
		return string(output), err
	}
}

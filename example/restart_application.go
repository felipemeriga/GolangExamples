package example

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func RestartExample() {
	t1 := time.NewTimer(time.Second * 10)
	wg.Add(1)
	restartChan := make(chan bool)

	// you can handle to restart programmatically
	go func() {
		defer wg.Done()
		<-restartChan
		RestartSelf()
	}()

	go func() {
		for {
			select {
			case <-t1.C:
				fmt.Println("Timer expired")
				restartChan <- true
				return
			default:
				// do something
			}
		}
	}()

	fmt.Println(time.Now().Format("2006-Jan-02 ( 15:04:05)"))
	rand.Seed(time.Now().UnixNano())
	// you can handle the shut-down programmatically
	//if rand.Intn(3) == 1 {
	//	fmt.Println("It is time to shut-down")
	//	os.Exit(0)
	//}
	wg.Wait()

}
func RestartSelf() error {
	self, err := os.Executable()
	if err != nil {
		return err
	}
	args := os.Args
	env := os.Environ()
	// Windows does not support exec syscall.
	if runtime.GOOS == "darwin" {
		cmd := exec.Command(self, args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Env = env
		err := cmd.Run()
		if err == nil {
			os.Exit(0)
		}
		return err
	}
	return syscall.Exec(self, args, env)
}

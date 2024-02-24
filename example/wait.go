package example

import (
	"fmt"
	"sync"
	"time"
)

func SimpleWaitingGroup() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		time.Sleep(time.Second * 2)
		fmt.Println("2 SECONDS!!")
	}()

	waitGroup.Wait()
	fmt.Println("finished!!")

}

func WaitGroupExample() {
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		time.Sleep(time.Second * 2)
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		time.Sleep(time.Second * 4)
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		time.Sleep(time.Second * 6)
	}()

	waitGroup.Wait()

	fmt.Println("Finished!")
}

package example

import (
	"fmt"
	"time"
)

func WaitingTicker() {
	ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("ae")
		}
	}
}

func processParallel(number int) {
	if number == 2 || number == 3 {
		time.Sleep(time.Second * 3)
	} else {
		time.Sleep(time.Second * 1)
	}

	<-sem
}

var sem = make(chan int, 3)

func Parallel() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, number := range numbers {
		sem <- number
		number := number
		go func() {
			processParallel(number)
		}()
	}
}

func ExecuteChannel() chan int {
	m := make(chan int)

	go func() {
		time.Sleep(6 * time.Second)
		m <- 5

	}()

	return m
}

func WaitingChannel() {
	//myChannel := ExecuteChannel()
	//
	//number := <- myChannel
	//
	//fmt.Println(number)
	//go func() {
	//	myChannel <- 6
	//	time.Sleep(2 * time.Second)
	//	myChannel <- 7
	//}()
	//
	//fmt.Println(myChannel)

	ch1 := make(chan string, 1)
	go func() {
		ch1 <- "Some value"
		time.Sleep(time.Second * 5)
		ch1 <- "Some value"
	}()

	go func() {
		for {
			<-ch1
			fmt.Println("brinbi")
		}
	}()

	time.Sleep(time.Second * 10)

}

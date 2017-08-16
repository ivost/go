package main

import (
	"fmt"
	"time"
)

func main2() {

	//fmt.Println(<- waitAndSend(42, 2))
	select {
		case v1 := <- waitAndSend(3, 1):
			fmt.Println(v1)
		case v2 := <- waitAndSend(5, 2):
			fmt.Println(v2)
		default:
			fmt.Println("DEFAULT")
	}
	time.Sleep(2500 * time.Millisecond)

}

// buffered channels block when they are either full (on write) or empty (on read)
func buferred() {
	//fmt.Println("Buffered channel demo")
	buffch := make(chan int, 3)
	buffch <- 1
	buffch <- 2
	buffch <- 3
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)
	fmt.Println(<-buffch)
}

// will wait for waitsec before sending value val on return channel
func waitAndSend(val, waitsec int) chan int {
	fmt.Println("waitAndSend enter")
	retCh := make(chan int)
	go func() {
		time.Sleep(time.Duration(waitsec) * time.Second)
		retCh <- val
	}()
	return retCh

}
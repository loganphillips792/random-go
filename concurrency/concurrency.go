package concurrency

import (
	"fmt"
	"time"
	"runtime"
	"sync"
)

// https://github.com/luk4z7/go-concurrency-guide
// https://www.youtube.com/watch?v=oV9rvDllKEg
// https://gobyexample.com/goroutines
// https://gobyexample.com/channels
// https://www.youtube.com/watch?v=VkGQFFl66X4
// https://www.youtube.com/watch?v=LvgVSSpwND8

func RunConcurrency() {
	fmt.Println("Running concurrency")
	
	// runBasicConcurrency()
	// runWaitGroupConcurrency()
	// runWithChannel() // blocking nature of channels allows us to sychronize goroutines
	// runWithChannelLoopDeadlock()
	runWithChannelLoopNoDeadlock()
}

func runBasicConcurrency() {

	/*
	This will end immediately because the main goroutine will have no more work to do since
	the 2 below functions are their own goroutines
	*/

	go count("sheep")
	go count("fish")
	fmt.Println(runtime.NumGoroutine()) // 3 goroutines. The main one + 2
}

func runWaitGroupConcurrency() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait() // wait until the counter is 0
}

func runWithChannel() {
	c := make(chan string)
	go countWithChannel("sheep", c)

	// blocking operation: when you try to receive something, you have to wait for the value
	msg := <- c
	fmt.Println(msg)
}

func runWithChannelLoopDeadlock() {
	c := make(chan string)
	go countWithChannel("sheep", c)

	for { // this causes a deadlock because once the countWithChannel() loop finishes, there aren't any more goroutines that will send to this and so it would sit forever. Go is able to recgonize this at run time.
		msg := <- c
		fmt.Println(msg)
	}
}

func runWithChannelLoopNoDeadlock() {
	c := make(chan string)
	go countWithChannelSolveDeadlock("sheep", c)

	/*
	for {
		msg, open := <- c // if you are a receiver, don't close a channel because you don't know if the sender is finished or not. The sender should close the channel
		
		if !open {
			break
		}
		
		fmt.Println(msg)
	}
	*/
	for msg := range c { // keep receiving messages until the channel is closed
		fmt.Println(msg)
	}
}


func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countWithChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// blocking operation: when you are sending a message, it will wait for a receiver
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
}

func countWithChannelSolveDeadlock(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// blocking operation: when you are sending a message, it will wait for a receiver
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
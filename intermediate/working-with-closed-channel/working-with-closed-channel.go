package main

import (
	"fmt"
	"time"
)

// Receiving from a closed channel is safe, whereas writing on closed channel throws a panic.
// Second value received from the channel, indicates that is there more moves to be received or not.
// This is a well-documented behavior, but it's not very intuitive for new Go developers
// who might expect the send behavior to be similar to the receive behavior.
// This is complex enough and needs quite a good amount of thought as this may
// be resolved with minor code change or may need change in design.
func main() {
	// trap()
	alternative()
}

func trap() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go func(a int) {
			ch <- (a + 1) * 2 // panic: send on closed channel
		}(i)
	}

	fmt.Println(<-ch) // get the first result
	close(ch)         // not ok (you still have other senders)
	time.Sleep(2 * time.Second)
}

func alternative() {
	ch := make(chan int)
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(a int) {
			select {
			case ch <- (a + 1) * 2:
				fmt.Println(a, "sent result")
			case <-done:
				fmt.Println(a, "exiting")
			}
		}(i)
	}

	fmt.Println("result:", <-ch) // get first result
	close(done)
	time.Sleep(3 * time.Second)
}

// Output :
// 10
// panic: send on closed channel
// 4 sent result
// result: 10
// 0 exiting
// 1 exiting
// 2 exiting
// 3 exiting

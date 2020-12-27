package main

import (
	"fmt"
	"time"
)

// You might think the sender will not be blocked until your message is processed by the receiver.
// Depending on the machine where you are running the code,
// the receiver goroutine may or may not have enough time to process the message
// before the sender continues its execution.
func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
			time.Sleep(time.Second)
		}
	}()

	ch <- "1"
	fmt.Println("Sending 2")
	ch <- "2" // Wait till reading of first value
	fmt.Println("Sending 3")
	ch <- "3" // Wait till reading of second value
	fmt.Println("Done Sending !!")
}

// Output :
// processed: 1
// Sending 2
// processed: 2
// Sending 3
// processed: 3
// Done Sending !!

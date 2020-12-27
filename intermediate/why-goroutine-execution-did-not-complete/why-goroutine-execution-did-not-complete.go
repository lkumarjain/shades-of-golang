package main

import (
	"fmt"
	"sync"
	"time"
)

// Application does not wait for all goroutines to complete before it exists.
// This is a most common mistake done by a developer in early days of learning Go.
// We can avoid this using WaitGroup a most common solution,
// it allows us to wait in the main for goroutines to complete.
// Another solution is to pass goroutine execution state using channel.
func main() {
	trap()        // Exists without executing all Routines
	alternative() // Wait for all routines to complete
}

func work(caller string, workerID int) {
	fmt.Printf("[%v] %s is running\n", workerID, caller)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] %s is done\n", workerID, caller)
}

func trap() {
	for i := 0; i < 5; i++ {
		go work("Trap", i) // Executing a function in routine
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Trap all done !!") // Prints: all done !!
}

func alternative() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1) // Adding a worker
		go func(i int) {
			defer wg.Done()        // Marking a working completion
			work("Alternative", i) // Executing a function in routine
		}(i)
	}
	wg.Wait() // Wait for all routines to complete
	fmt.Println("Alternative all done !!")
}

// Output :
// [4] Trap is running
// [0] Trap is running
// [1] Trap is running
// [2] Trap is running
// [3] Trap is running
// Trap all done !!
// [4] Alternative is running
// [2] Alternative is running
// [3] Alternative is running
// [1] Alternative is running
// [0] Alternative is running
// [0] Trap is done
// [4] Trap is done
// [1] Trap is done
// [2] Trap is done
// [3] Trap is done
// [1] Alternative is done
// [2] Alternative is done
// [0] Alternative is done
// [3] Alternative is done
// [4] Alternative is done
// Alternative all done !!

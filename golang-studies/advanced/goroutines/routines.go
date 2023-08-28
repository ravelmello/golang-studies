package routines

import (
	"fmt"
	"time"
)

func Letters(done chan<- bool) {
	for l := 'a'; l <= 'z'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(100 * time.Millisecond)
	}
	done <- true
}

func Numbers(number chan<- int) {
	for i := 0; i <= 23; i++ {
		number <- i
		fmt.Printf("Write in channel %d\n", i)
	}

	close(number)

	fmt.Printf("End of writing in channel.\n")
}

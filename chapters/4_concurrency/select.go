package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick: // Try to read from channel 'tick'
			fmt.Println("tick.")
		case <-boom: // Try to read from channel 'boom'
			fmt.Println("BOOM!")
			return
		default: // If no other cases are ready
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

package main

import (
	"fmt"	
	"time"
)

func printWithDelay(i int) {
	fmt.Printf("iteration number: %d \n", i)
	
}

func main() {
	for i := 0; i < 5; i++ {
		go printWithDelay(i)	
	}
	time.Sleep(5 * time.Second)
}



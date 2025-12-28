package main

import (
	"fmt"
)

func printWithDelay(i int) {
	fmt.Printf("iteration number: %d \n", i)
	
}

func addToChannel(c chan int) {

	for i := 10; i < 15; i++ {
		c <- i
	}
	return 	
}

func main() {

	testChannel := make(chan int)

//	for i := 0; i < 5; i++ {
//		go printWithDelay(i)	
//	}

	go addToChannel(testChannel)

	for i := 0; i < 5; i++ {
		
		val, ok := <-testChannel
		if !ok {
			fmt.Printf("No more numbers\n")
			break
		}
		fmt.Printf("current number is : %d \n", val)
	}
}




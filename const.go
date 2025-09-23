package main

import "fmt"

func main() {

	const minutesInAnHour = 60
	const secondsInAMinute = 60
	const secondsInAnHour = 60*60 // Constants can be a computed value but must be known at compile time

	mutableSecondsInAnHour = 60*60

	fmt.Println("Number of seconds in an hour:", mutableSecondsInAnHour)

}

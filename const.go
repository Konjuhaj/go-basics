package main

import "fmt"
import "time"

func Const() {

	const minutesInAnHour = 60
	const secondsInAMinute = 60
	const secondsInAnHour = 60*60 // Constants can be a computed value but must be known at compile time

	mutableSecondsInAnHour := 60*60 // Var with walrus operator and computed value

	fmt.Println("Number of seconds in an hour:", mutableSecondsInAnHour)

	// Does not work due to time being calculated at run time
	// const currentTime = time.Now() 
	currentTime := time.Now()
	fmt.Println(currentTime)
}

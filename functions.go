package main

import "fmt" 

func ignorePram() (int, string, string) {

	return 18, "Besnik", "Konjuhaj"

}

func main() {
	fmt.Println("Stargin Unit Tests") 

	age, firstName, _ := ignorePram()

	fmt.Printf("Hello my name is %v and I'm %d years old", firstName, age) 

}

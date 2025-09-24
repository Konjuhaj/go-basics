package main

import "fmt"

func Conditions() {
	amountOfPeople := 2221
	const venueLimit int = 30

	if amountOfPeople <= venueLimit {
		fmt.Println("Everyone got in")
	} else {
		fmt.Printf("21 people got in and %v were left out \n",  amountOfPeople - venueLimit)
	}
	if outsiders := amountOfPeople - venueLimit; outsiders > 0 {
		fmt.Printf("%d people went home and are sad\n", outsiders)	
	}

}

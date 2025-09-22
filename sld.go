package main

import "fmt"

func Sld() {

	one, two, three := 1, 2, 3

	fmt.Println("Same line declaration with same type")
	fmt.Println(one, two, three)
	
	fmt.Println("Same line declaration with multi types")

	s, b, f := "a string", true, 0.21
	fmt.Println(s, b, f)	

	// Three variables, two assignments
	//a, b := 1, 2, 3
	//fmt.Println(a, b, c)
	// ERROR MESSAGE: ./sld.go:18:13: assignment mismatch: 3 variables but 2 values

	
}

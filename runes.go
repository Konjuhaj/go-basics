package main

import "fmt"
import "unicode/utf8"

func main() {

	const name = "test" // A string with 4 characters and four runes
	const bear = "🐻" // an emoji with four characters and one rune
	// Runes are of size 32 bits and character are of size 8 bits -- I.e standard ascii

	fmt.Printf("There are %d characters in name \n", len(name))
	fmt.Printf("There are %d runes and %d characters in 🐻\n", utf8.RuneCountInString(bear), len(bear))

}

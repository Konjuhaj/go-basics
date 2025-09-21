package main

import "fmt"

func main() {
	mySkillIssue := 42
	pi := 3.14159
	message := "hello world"
	isGoat := true

	fmt.Printf("%v, %v, %v, %v\n", mySkillIssue, pi, message, isGoat)
	fmt.Println(mySkillIssue, pi, message, isGoat)
	fmt.Printf("%T, %T ,%T, %T", mySkillIssue, pi, message, isGoat)
}

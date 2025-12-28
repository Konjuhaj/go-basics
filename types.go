package main 

import "fmt"

type User struct{
	username string
	userId	int
}

func updateUserId(u *User) {
	u.userId = 444
	fmt.Printf("userId in function is %d \n", u.userId)
}

func (u User) introduce(greeting string) {
	fmt.Printf("%s %s", greeting, u.username)
} 

func main() {
	user := User{"besnik", 123}
	defer user.introduce("Hello my username is")	

	updateUserId(&user)
	fmt.Printf("%s\n", user.username)
	fmt.Printf("UserID in main is %d\n", user.userId)
}

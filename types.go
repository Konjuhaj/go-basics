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


func main() {
	user := User{"besnik", 123}
	updateUserId(&user)
	fmt.Printf("%s\n", user.username)
	fmt.Printf("UserID in main is %d\n", user.userId)
}

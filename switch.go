package main 

import "fmt"

func billingCost(plan string){
	switch plan {
		case "basic":
			return 10.00
		default:
			return 0.0
	}
}

func main() {
	 
	plan := "basic"
	fmt.Println("You're current plan costs: ", billingCost(plan))

	fmt.Println("You're current plan costs: ", billingCost("test"))
}

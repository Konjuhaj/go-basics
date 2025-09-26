package main 

import "fmt"

func billingCost(plan string) float64{


	switch plan {
		case "basic":
			return 10.00		
		case "premium":
			fallthrough
		case "deluxe":
			return 20.00
		default:
			return 0.0
	}
}

func main() {
	 
	plan := "basic"
	fmt.Println("You're current plan costs: ", billingCost(plan))

	fmt.Println("You're current plan costs: ", billingCost("test"))
}

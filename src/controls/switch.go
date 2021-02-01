package main

import "fmt"

func main() {
	city := "Madurai"

	switch city {
	case "Chennai", "Madurai":
		fmt.Println("You live in Tamilnadu")
	case "Bangalore":
		fmt.Println("You live in Karnataka")
	case "Trivandrum":
		fmt.Println("You live in Kerala")
	default:
		fmt.Println("You're not from around here")
	}

	index := 100

	switch {
	case index > 10:
		fmt.Println("Greater than 10")
	case index < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Is 10")

	}

	index = 9
	switch {
	case index != 10:
		fmt.Println("Does not equal 10")
		fallthrough
	case index < 10:
		fmt.Println("Less than 10")
	case index > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("Is 10")
	}
}

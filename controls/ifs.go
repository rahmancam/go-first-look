package main

import "fmt"

func main() {
	var someVar = 9

	if someVar > 10 {
		fmt.Println(someVar)
	}

	if someVar > 100 {
		fmt.Println("Greater than 100")
	} else if someVar == 100 {
		fmt.Println("Equals 100")
	} else {
		fmt.Println("Less than 100")
	}

	// err := someFunction()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if err := someFunction(); err != nil {
	// 	fmt.Println(err.Error())
	// }

}

// func someFunction()  {
// 	return 0
// }

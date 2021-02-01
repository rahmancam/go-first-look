package main

import "fmt"

func main() {
	var scores [5]float64
	fmt.Println(scores)
	fmt.Println(len(scores))
	scores[0] = 9
	scores[1] = 1.5
	scores[2] = 4.5
	scores[3] = 7
	scores[4] = 8

	var scores2 [5]float64 = [5]float64{9, 1.5, 4.5, 7, 8}
	fmt.Println(scores2)

	scores3 := [5]float64{9, 1.5, 4.5, 7, 8}
	fmt.Println(scores3)

	scores4 := [...]float64{9, 1.5, 4.5, 7, 8, 10, 12.6, 18.2}
	fmt.Println(scores4)
	fmt.Println(len(scores4))

	for _, value := range scores4 {
		fmt.Println(value)
	}
}

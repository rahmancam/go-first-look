package main

import "fmt"

func main() {
	age, after40 := getAge(8)
	fmt.Println(age, after40)
	age1, age2 := getAge2()
	fmt.Println(age1, age2)
	printAge(16, 21, 33)
	printAge2(16, 21, 33, 45, 96, 101)
	fmt.Println(average(10, 20, 30))
	fmt.Println(mean(10, 20, 30))
}

func getAge(age int) (int, int) {
	return age, age + 40
}

func getAge2() (ageOfSally int, ageOfBob int) {
	ageOfBob = 21
	ageOfSally = 16
	return
}

func printAge(age1, age2, age3 int) {
	fmt.Println(age2, age2, age3)
}

func printAge2(ages ...int) {
	for _, val := range ages {
		fmt.Println(val)
	}
}

func average(num1, num2, num3 float64) float64 {
	total := num1 + num2 + num3
	return total / 3
}

func mean(nums ...float64) float64 {
	total := 0.0
	for _, number := range nums {
		total += number
	}
	return total / float64(len(nums))
}

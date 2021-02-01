package main

import "fmt"

func average(scores [5]float64) float64 {
	total := 0.0
	for _, num := range scores {
		total += num
	}
	return total / float64(len(scores))
}

var pets map[string]string = map[string]string{
	"fido":     "dog",
	"penelope": "horse",
	"nancy":    "cat",
}

func doesPetExist(petName string) bool {
	_, ok := pets[petName]
	return ok
}

var groceries = []string{"beans", "lemons", "chicken", "fruit"}

func addGroceriesToList(newGroceries ...string) []string {
	foods := groceries
	for _, g := range newGroceries {
		foods = append(foods, g)
	}
	return foods
}

func main() {
	scores := [5]float64{2, 7, 8, 1, 9}
	fmt.Println(average(scores))
	fmt.Println("Does pet exist?", doesPetExist("fido"))

	groceryList := addGroceriesToList("beets", "choclate")
	fmt.Println(groceryList)
}

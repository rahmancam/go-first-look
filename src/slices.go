package main

import "fmt"

func main() {
	var myArray [5]int
	var mySlice []int = make([]int, 5, 10)

	myArray[0] = 1
	mySlice[0] = 1

	fmt.Println(myArray)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	fruitArray := [5]string{"banana", "pear", "apple", "peach", "orange"}
	splicedFruit := fruitArray[1:3]
	fmt.Println(splicedFruit)
	fmt.Println(len(splicedFruit))
	fmt.Println(cap(splicedFruit))

	fruitToAdd := append(splicedFruit, "canteloupe", "cherries")

	fmt.Println(splicedFruit, fruitToAdd)
	fmt.Println(len(splicedFruit), cap(splicedFruit))
	fmt.Println( len(fruitToAdd), cap(fruitToAdd))


	fruitToAdd2 := append(splicedFruit, "canteloupe", "cherries", "lemon")

	fmt.Println(splicedFruit, fruitToAdd2)
	fmt.Println(len(splicedFruit), cap(splicedFruit))
	fmt.Println( len(fruitToAdd2), cap(fruitToAdd2))

}

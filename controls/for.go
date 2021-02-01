package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 100 {
		fmt.Println(i)
		i++
	}

	mySentence := "This is a sentence"

	for index, letter := range mySentence {
		fmt.Println("Index:", index, "Letter:", string(letter))
	}

	for _, letter := range mySentence {
		fmt.Println(string(letter))
	}

	for index, letter := range mySentence {
		if index%2 != 0 {
			fmt.Println("Index: ", index, "Letter:", string(letter))
		}
	}
}

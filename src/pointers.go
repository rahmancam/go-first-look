package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	var namePointer *string

	name = "Abdul"
	namePointer = &name

	var nameValue = *namePointer

	fmt.Println("Name:", name)
	fmt.Println("Name *:", namePointer)
	fmt.Println("Name Value:", nameValue)

	changeName(&name)
	fmt.Println(name)

	var pt = Coordinates{X: 10, Y: 20}
	fmt.Println(pt)
	updateCoords(&pt)
	fmt.Println(pt)
}

func updateCoords(pt *Coordinates) {
	pt.X = 200
}

func changeName(n *string) {
	*n = strings.ToUpper(*n)
}

// Coordinates X,Y
type Coordinates struct {
	X, Y float64
}

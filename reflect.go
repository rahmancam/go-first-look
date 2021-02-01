package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf("Abdul Rahman"))
	var x = 4
	fmt.Println(reflect.TypeOf(float64(x) * 5.5))
}

package main

import (
	"errors"
	"fmt"
	"os"
)

func isGreaterThanTen(num int) error {
	if num < 10 {
		return errors.New("Something bad happened")
	}
	return nil
}

func main() {
	num := 9
	if err := isGreaterThanTen(num); err != nil {
		fmt.Println(fmt.Errorf("%d is NOT GREATER THAN TEN", num))
		// panic(err)
		// log.Fatalln(err)
	}

	if err := openFile(); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	doThings()
	doThings2()
}

func openFile() error {
	f, err := os.Open("missingFile.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func doThings() {
	defer fmt.Println("First line but do this last!")
	defer fmt.Println("Do this second to last")
	fmt.Println("Things and stuff shluld happen first")
}

func doThings2() {
	defer recoverFromPanic()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			panic("PANIC!")
		}
	}
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("We panicked but evreyones's fine")
		fmt.Println(r)
	}
}

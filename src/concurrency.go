package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC")
	}
}

func say(s string) {
	defer wg.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("There")
	wg.Wait()
}

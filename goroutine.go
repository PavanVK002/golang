package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sayHello() // Run sayHello concurrently as a goroutine
	time.Sleep(500 * time.Millisecond) // Wait for a while to allow goroutine to finish
	fmt.Println("Main function")
}

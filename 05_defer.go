package main

import "fmt"

func main() {
	// defer can be an array of callbacks
	defer fmt.Println("world")

	fmt.Println("hello")
}

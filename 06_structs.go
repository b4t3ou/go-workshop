package main

import "fmt"

type Coordinates struct {
	X float32
	Y float32
}

func main() {
	// Array
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Slice
	var slice []int = primes[1:4]
	fmt.Println(slice)

	t := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len = %d capacity = %d %v\n", len(t), cap(t), t)

	slice2 := make([]int, 15)
	//slice2 := make([]int, 15, 200)
	fmt.Printf("len = %d capacity = %d %v\n", len(slice2), cap(slice2), slice2)

	// Struct
	v := Coordinates{1, 2}
	p := &v
	p.X = 100
	fmt.Println(v)

	v2 := &Coordinates{1, 2}
	v2.Y = 200
	fmt.Println(*v2)

	// Maps
	m := make(map[string]Coordinates)
	m["London"] = Coordinates{40.68433, -74.39967}
	fmt.Println(m, m["London"])
}

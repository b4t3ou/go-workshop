package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)


	var t int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", t, f, b, s)
}
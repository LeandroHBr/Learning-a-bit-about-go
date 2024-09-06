package main

import "fmt"

func main() {

	x := 3
	y := 3

	if x > y {
		fmt.Println("X > Y")
	} else if x < y {
		fmt.Println("X < Y")
	} else {
		fmt.Println("X = Y")
	}
}

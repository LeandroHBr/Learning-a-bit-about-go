package main

import "fmt"

func main() {
	var names = [3]string{"carlos", "ronaldo", "pedro"}
	var numbers = [...]float32{1.1, 2, 3, 55.5}
	fmt.Println("names:", names)
	fmt.Println("numeber of index 1:", numbers[1])
	numbers[1] = 9
	fmt.Println("new number of index 1:", numbers[1])
	fmt.Println("the length of array 'names' is", len(names))
}

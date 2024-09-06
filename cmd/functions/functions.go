package main

import "fmt"

/*func greet(name string, age int) {
	fmt.Println("hi my name is", name, "and i have ", age, "year old")
}*/

func sum(a int, b int) int {
	return a + b
}

func main() {
	/*
		greet("Leandro", 18)
		greet("hélio", 29)*/

	var c int = sum(1, 2)
	fmt.Println(c)

}

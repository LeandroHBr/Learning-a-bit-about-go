package main

import "fmt"

/*
	1.int- stores integers (whole numbers), such as 123 or -123
	2.float32- stores floating point numbers, such as 19.99.
	3.string - stores text, such as "Hello World".
	4.bool- stores values with two states: true or false

	syntax -> var varible type = value
					or
				var variable := value
					or
					var variable = value
*/

func main() {
	var name string = "leandro"
	var age int = 18
	fmt.Println("Hi, my name is", name, "and i have", age, "years old")
}

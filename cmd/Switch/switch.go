package main

import "fmt"

func main() {
	name := "leandro"

	switch name {
	case "pedro":
		fmt.Println("hi pedro")
	case "carlos":
		fmt.Println("hi, carlos")
	case "leandro":
		fmt.Println("hi, leandro")
	default:
		fmt.Println("hi")
	}
}

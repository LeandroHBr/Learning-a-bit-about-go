package main

import "fmt"

type Person struct {
	Name         string
	age          int
	job          string
	phone_number string
	genre        string
}

func main() {

	Person1 := Person{"leandro", 18, "programmer", "12345678", "male"}

	
	fmt.Println("name: ", Person1.Name, "\ngenre: ", Person1.genre)
}

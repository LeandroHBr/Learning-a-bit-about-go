package main

import "fmt"

func main() {
	var car = map[string]string{
		"brand": "Toyota",
		"year":  "2006",
		"price": "$2000",
	}

	for i, value := range car {
		fmt.Println(i, value)
	}
}

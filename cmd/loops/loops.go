package main

func main() {
	/*
		for i := 0; i < 10; i++ {
			if i == 3 {
				continue -> jump in i == 3
			}
			if i > 8 {
				break
			}
			println(i)
		}
	*/

	var cars = [5]string{"gol", "toyota", "gipe", "uno", "cross_fox"}

	for i := 0; i < len(cars); i++ {
		println("car", i+1, ":", cars[i])
	}

}

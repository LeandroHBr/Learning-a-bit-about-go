package main

import "fmt"

/*slice is similar to an array, but must powerfull
declaration -> name := []type{values}


capacity vs lenght.
capacity is how many alements the slice can 'store'.
len is how many elements the slice have
*/

func main() {
	myslice := []int{1, 2, 3, 4}
	fmt.Println(myslice)
	fmt.Println("slice capacity:", cap(myslice))
	fmt.Println("slice lenght:", len(myslice))

	// appending 2 slices

	myslice2 := make([]int, 5, 10)
	copy(myslice2, myslice)

	myslice3 := append(myslice, myslice2...)
	fmt.Println("new slice", myslice3)

}

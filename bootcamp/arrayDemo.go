package main

import (
	"fmt"
)

func main() {
	singleArray()
	fmt.Print("\n")
	multidimArray()

}

func singleArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// another way to declare array
	b := [2]string{"Hello", "World!"}
	fmt.Printf("%q", b)

	/*
		use an ellipsis to use an implicit length
	*/
	fmt.Println("use an ellipsis to use an implicit length")
	c := [...]string{"Hello", "World!"}
	fmt.Printf("%q", c)
}

func multidimArray() {
	var a [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = fmt.Sprintf("row %d - column %d", i+1, j+1)
		}
	}
	// %q to print array with structure
	fmt.Printf("%q\n", a)
}

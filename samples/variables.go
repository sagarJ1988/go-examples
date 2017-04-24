package main

import "fmt"

func main() {
	// variable takes data type from right hand side value
	var a string = "intial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// default value must be 0
	var e int
	fmt.Println(e)

	// short hand declaration and intialization of variable
	f := "short"
	fmt.Println(f)
}

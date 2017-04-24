package main

import "fmt"

func main() {
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World\n")
	// boolean demos
	fmt.Println("boolean demo\n")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// varaible demo
	fmt.Println("varaible demo \n")
	var x string
	x = "first"
	fmt.Println(x)
	x += " second"
	fmt.Println(x)

	//
	y := 5
	fmt.Println(y)
}

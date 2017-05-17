package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("variable declaration")
	variablesDec()
	fmt.Println("\nconstant declaration")
	constants()
	fmt.Println("variables and constant print")
	varConstPrint()
}

func variablesDec() {
	/*
		variable declaration
		1)
			var (
				name     string
				age      int
				location string
			)
		2)
			var (
				name, location string
				age            int
			)
		3)
			var name string
			var age int
			var location string
		4)
			var (
					name string = "Prince Oberyn"
					age int = 32
					location string = "Dorne"
				)
		5)
			var (
					name = "Prince Oberyn"
					age = 32
					location = "Dorne"
				)
		6)
			var (
				name, location, age = "Prince Oberyn", "Dorne", 32
			)

		:=  only use in function
	*/

	name, location := "Prince Oberyn", "Dorne"
	age := 32
	fmt.Printf("%s (%d) of %s", name, age, location)
}

func constants() {
	/*
		constants are immutable
	*/
	const (
		Pi    = 3.14
		Truth = false
		Big   = 1 << 62
		Small = Big >> 61
	)

	const Greeting = ""
	fmt.Println(Greeting)
	fmt.Println(Pi)
	fmt.Println(Truth)
	fmt.Println(Big)

}

func varConstPrint() {
	name := "Caprica-Six"
	aka := fmt.Sprintf("Number %d ", 6)
	fmt.Printf("%s is also known as %s\n",
		name, aka)

	fmt.Println("math", math.Pi)
}

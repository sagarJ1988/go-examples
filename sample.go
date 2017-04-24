package main

import "fmt"

// this is a comment
func main() {
	var x string
	x = "Hello World"
	fmt.Println(x)
	x = "Hello"
	fmt.Println(x)
	// fmt.Println(len("Hello World"))
	// fmt.Println("Hello World"[10])
	// fmt.Println("Hello " + "World")
	var n = 10
	fiboSeries(n)
	primeNumber()
}

func fiboSeries(n int) {
	//fmt.Println("test")
	a, b := 0, 1
	fmt.Println("fiboSeries:")
	//fmt.Println(a)
	//	fmt.Println(b)

	for i := 0; i <= 10; i++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func primeNumber() {
	var i int
	fmt.Println("primeNumber:")
	for i <= 10 {
		if i%2 == 0 {
			fmt.Printf("Even %d", i)
		} else {
			fmt.Printf("Odd %d", i)
		}
		fmt.Printf("\n")
		i++
	}
}

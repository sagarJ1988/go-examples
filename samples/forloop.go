package main

import "fmt"

func main() {
	i := 1
	fmt.Println("first for loop")
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("second for loop")
	for j := 0; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("third for loop")
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("fourth for loop")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

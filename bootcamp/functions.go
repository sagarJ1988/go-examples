package main

import (
	"fmt"
)

// declare multiple type
func add(x int, y int) int {
	return x + y
}

// declare single type
func add1(x, y int) int {
	return x + y
}

//return multiple parameters
func location(city string) (string, string) {
	var region string
	var continent string

	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region, continent = "California", "North America"
	case "New York", "NYC":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

// return named parameters
func location1(name, city string) (region, continent string) {
	switch city {
	case "New York", "LA", "Chicago":
		region, continent = "New York", "North America"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return
}

func main() {
	fmt.Println("simple function with return single value declare multiple type")
	fmt.Println(add(42, 13))

	fmt.Println("simple function with return single value declare single type")
	fmt.Println(add1(42, 13))

	fmt.Println("simple function with return multiple value")
	region, continent := location("Santa Monica")
	fmt.Printf("ABC lives in %s, %s\n", region, continent)

	fmt.Println("simple function with return named value")
	region, continent = location1("ABC", "LA")
	fmt.Printf("ABC lives in %s, %s \n", region, continent)
}

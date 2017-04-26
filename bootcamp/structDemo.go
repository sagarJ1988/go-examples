package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	Lat, Lon float64
	Date     time.Time
}

type Bootcamp1 struct {
	Lat, Lon float64
}

func main() {
	// intialization
	event := Bootcamp{
		Lat: 34.012836,
		Lon: -118.495338,
	}

	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f,%f)\n", event.Date, event.Lat, event.Lon)

	/*Go supports the new expression to allocate a zeroed value of the requested type
	and to return a pointer to it
	*/
	fmt.Println("with new ")
	x := new(Bootcamp1)
	y := &Bootcamp1{}
	fmt.Println(*x == *y)
}

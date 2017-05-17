package main

import (
	"fmt"
)

type Artist struct {
	Name, Genre string
	Songs       int
}

// function pass by value as parameter
func newRelease(a Artist) int {
	a.Songs++
	return a.Songs
}

func newRelease1(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	/*
		pass by value
		just passing struct value
		In pass by value original value not changed
		output:
		ABC released their 43th song
		ABC has a total of 42 songs
	*/
	me := Artist{Name: "ABC", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs\n", me.Name, me.Songs)

	/*
		pass by reference
		In pass by reference original value changed
		output:
		ABC released their 43th song
		ABC has a total of 43 songs
	*/

	me1 := &Artist{Name: "ABC", Genre: "Electro", Songs: 42}
	fmt.Printf("%s released their %dth song\n", me1.Name, newRelease1(me1))
	fmt.Printf("%s has a total of %d songs\n", me1.Name, me1.Songs)
}

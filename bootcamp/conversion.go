package main

import (
	"fmt"
	"time"
)

/*
	Type conversion

	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	or

	i := 42
	f := float64(i)
	u := uint(f)
*/

// type assertion
/*
	If you have a value and want to convert it to another or a specific type (in case
of interface{} ), you can use type assertion
*/

type Sample interface {
	String() string
}

type Test struct {
	content string
}

// we pass empty interface
func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

//function used to implement sample interface

func (s *Test) String() string {
	return s.content
}

func displayString(val interface{}) {
	switch str := val.(type) {
	case string:
		fmt.Println(str)
	case Sample:
		fmt.Println(str.String())
	}
}

func main() {
	/***** not an assertion*********/
	foo := map[string]interface{}{
		"ABC": 42,
	}

	timeMap(foo)
	fmt.Println(foo)
	/*******not an assertion********
	output
	map[ABC:42 updated_at:2017-04-24 19:22:27.932898752 +0530 IST]
	*/

	/*******assertion*******
		when you have a function taking a param of a specific interface but the
	function inner code behaves differently based on the actual object type
	output
	Assertion Demo
	Assertion string
	*/

	s := &Test{"Assertion Demo"}
	displayString(s)
	displayString("Assertion string")
}

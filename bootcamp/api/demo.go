package api

import (
	"fmt"
	"go-examples/sample-program/rest"
)

type Player struct {
	Name string
	Age  int
}

func Demo() {
	p := Player{}
	p.Age = 42
	p.Name = "Matt"
	fmt.Println("Inside Api package")
	fmt.Printf("%+v", p)
	rest.Rest()

}

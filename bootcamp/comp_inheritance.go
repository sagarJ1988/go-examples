package main

import (
	"fmt"
)

type User struct {
	Id       int
	Name     string
	Location string
}

type Player struct {
	Id       int
	Name     string
	Location string
	GameId   int
}

/*
	Player struct has same fields as the User struct except GameId
	duplicate fields are Id, Name,Location
	so we can define it as follows:
*/

type UserComposition struct {
	Id             int
	Name, Location string
}

type PlayerComposition struct {
	UserComposition
	GameId int
}

/*
	methods on user struct are also availabel for PlayerComposition
*/
func (u *UserComposition) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

func main() {
	/*
		without compostion
	*/
	fmt.Println("without compostion")
	p := Player{}
	p.Id = 42
	p.Name = "ABC"
	p.Location = "Pune"
	p.GameId = 1234
	fmt.Printf("%+v\n", p)

	/*
		%v	the value in a default format
			when printing structs, the plus flag (%+v) adds field names
		%#v	a Go-syntax representation of the value
		%T	a Go-syntax representation of the type of the value
		%%	a literal percent sign; consumes no value
	*/

	/*
		with compostion
	*/
	fmt.Println("with compostion")
	pcomp := PlayerComposition{}
	pcomp.Id = 43
	pcomp.Name = "PQR"
	pcomp.Location = "Mumbai"
	pcomp.GameId = 7896
	fmt.Printf("%+v\n", pcomp)

	/*
		other option using struct literal
	*/
	fmt.Println("struct literal with composition")
	pc := PlayerComposition{
		UserComposition{Id: 42, Name: "Matt", Location: "LA"},
		90404,
	}
	fmt.Printf(
		"Id: %d, Name: %s, Location: %s, Game id: %d\n",
		pc.Id, pc.Name, pc.Location, pc.GameId)
	// Directly set a field defined on the Player struct
	pc.Id = 11
	fmt.Printf("%+v\n", pc)

	/*
		it call UserComposition struct method using PlayerComposition
	*/

	p1 := PlayerComposition{}
	p1.Id = 20
	p1.Name = "ABC"
	p1.Location = "Pune"
	fmt.Println(p1.Greetings())
}

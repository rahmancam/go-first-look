package main

import "fmt"

// User type
type User struct {
	ID                         int
	FirstName, LastName, Email string
	age                        int
}

// Group type
type Group struct {
	role           string
	users          []User
	newstUser      User
	spaceAvailable bool
}

func (u User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (g *Group) describe() string {

	if len(g.users) > 2 {
		g.spaceAvailable = false
	}

	desc := fmt.Sprintf("This user group has %d. The newest user is %s %s. Accepting New Users: %t", len(g.users), g.newstUser.FirstName, g.newstUser.LastName, g.spaceAvailable)
	return desc
}

func main() {
	usr := User{ID: 1, FirstName: "Abdul", Email: "rmn@mail.com", LastName: "Rahman", age: 30}
	fmt.Println(usr)
	fmt.Println(usr.FirstName)
	usr2 := User{ID: 1, FirstName: "John", Email: "john@mail.com", LastName: "Smith", age: 33}
	usr3 := User{ID: 1, FirstName: "Peter", Email: "pete@mail.com", LastName: "Parker", age: 40}

	g := Group{
		role:           "admin",
		users:          []User{usr, usr2, usr3},
		newstUser:      usr2,
		spaceAvailable: true,
	}

	fmt.Println(usr.describe())
	fmt.Println(usr2.describe())

	fmt.Println(g.describe())
	fmt.Println(usr.describe())
	fmt.Println(usr2.describe())
	fmt.Println(g.describe())
}

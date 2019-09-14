package main

import (
	"fmt"
	// here we are importing a package
	// a package is essentially a folder with go code.
	"github.com/Sher-Chowdhury/gsg_child_packages/members"
)

func main() {

	/*
		notice the . notation for calling something from the child folder.
		similar to the dot syntax used in fmt.Println(x)
	*/

	Charlie := members.Employee{
		ID:        23,
		Title:     "Mr",
		FirstName: "Charles",
		Lastname:  "Dickens",
		Age:       55,
		Smoker:    true,
	}

	fmt.Println(Charlie) // prints {23 Mr Charles Dickens 55 true}
}

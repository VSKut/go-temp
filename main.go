package main

import (
	"fmt"
	"github.com/vskut/go-temp/people"
)

func main() {

	var peopleList people.People

	if p, err := people.NewPerson("Ivan", "Ivanov", "2005-08-10"); err == nil {
		peopleList = peopleList.Append(p)
	}
	// Also we can use else if we need to process error data

	if p, err := people.NewPerson("Ivan", "Ivanov", "2003-08-05"); err == nil {
		peopleList = peopleList.Append(p)
	}

	if p, err := people.NewPerson("Artiom", "Ivanov", "2005-08-10"); err == nil {
		peopleList = peopleList.Append(p)
	}

	peopleList.Sort()

	fmt.Println(peopleList)
}

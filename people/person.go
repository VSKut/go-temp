package people

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

func NewPerson(firstName string, lastName string, birthDay string) (Person, error) {
	birthDayTimeFormat, err := time.Parse("2006-08-02", birthDay)
	if err != nil {
		return Person{}, fmt.Errorf("an error with %s %s. Error: %s", firstName, lastName, err)
	}

	return Person{
		firstName: firstName,
		lastName:  lastName,
		birthDay:  birthDayTimeFormat,
	}, nil
}

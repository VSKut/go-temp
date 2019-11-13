package main

import (
	"fmt"
)

type Object struct {
	Name string
}

func main() {
	object := Object{Name: "world"}

	fmt.Println(object.greeting())
}

func (o *Object) greeting() string {
	return fmt.Sprintf("Hello %s!", o.Name)
}

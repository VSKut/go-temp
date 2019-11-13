package main

import (
	"testing"
)

func TestGreetingFormat(t *testing.T) {
	object := Object{Name: "test"}

	if r := object.greeting(); r != "Hello test!" {
		t.Error("greeting() != \"Hello test!\". Returned:", r)
	}

}

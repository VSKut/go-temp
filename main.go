package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	_, err := emoji.Println("Hello world :smile:!")
	if err != nil {
		fmt.Println(err)
	}
}

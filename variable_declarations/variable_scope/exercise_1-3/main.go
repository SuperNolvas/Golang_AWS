package main

import (
	"fmt"
)

var i = 100

func main() {
	fmt.Println("Scope global:", i)
	var i = 200
	fmt.Println("Scope main:", i)
}

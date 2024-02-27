package main

import "fmt"

func main() {
	var names [5]string
	names[0] = "Phil"
	names[2] = "Joe"
	fmt.Println("My names are: " + names[0])
	fmt.Println("There is no one here: " + names[1])
	fmt.Println("Who are you?: " + names[2])
}

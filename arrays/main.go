// package main

// import "fmt"

// func main() {
// 	var names [5]string
// 	names[0] = "Phil"
// 	names[2] = "Joe"
// 	fmt.Println("My names are: " + names[0])
// 	fmt.Println("There is no one here: " + names[1])
// 	fmt.Println("Who are you?: " + names[2])
// }

package main

import "fmt"

func main() {
	names := [...]string{"hello", "how", "are"}
	var namePointer [3]*string // this is a pointer to a string
	namePointer[0] = &names[0]
	namePointer[1] = &names[2] // this is a pointer to the 3rd element. Position 0 and 1 of the array have values set. POsition 2 is nil as it has not been set.

	for i, aPointer := range namePointer { // range returns the index and the value of the array
		fmt.Println(i, " : ", aPointer) // this will print the index and the value of the array
	}
}

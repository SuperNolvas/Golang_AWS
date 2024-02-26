package main

func main() {
	coat := "I am a coat" // value of coat

	receipt := &coat  // pointer to coat
	receipt2 := &coat // pointer to coat

	// Printing the values
	println(*receipt)                // dereference pointer
	println(*receipt2)               // dereference pointer
	println("Value of coat >", coat) // value of coat
	println(receipt)                 // address of coat
	println(receipt2)                // address of coat
	println(&coat)                   // address of coat
	println(&receipt)                // address of pointer
	println(&receipt2)               // address of pointer
	println(receipt == receipt2)     // true
	println(receipt == &coat)        // true
	println(receipt2 == &coat)       // true
	// println(receipt == &receipt2)    false  invalid operation: receipt == &receipt2 (mismatched types *string and **string) receipt is a pointer to string, &receipt2 is a pointer to pointer to string
	var intPointer *int // pointer to int
	println(intPointer) // nil
	//println(*intPointer)              panic: runtime error: invalid memory address or nil pointer dereference
	//*intPointer = 10                  panic: runtime error: invalid memory address or nil pointer dereference
}

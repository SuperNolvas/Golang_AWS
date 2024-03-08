package main

import "fmt"

type Entry struct {
	EventVersion string
	EventSource  string
	Sns          struct { // nested struct
		Message string
	}
}

// SNS is a struct that contains a slice of Entry structs
type SNS struct {
	Records []Entry
}

func main() {
	var message SNS
	var entry Entry

	entry.EventSource = "aws:sns"
	entry.EventVersion = "1.0"
	entry.Sns.Message = "Hello from SNS"

	message.Records = make([]Entry, 1) // create slice of length 1
	message.Records[0] = entry

	fmt.Println(message)
}

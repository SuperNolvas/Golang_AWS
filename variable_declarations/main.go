package main

import "fmt"

func main() {
	var VarInt int = 42
	var varString = "Test"
	var varFloat string
	varFloat = "3.14"
	varBool := true

	fmt.Printf("Value Int : %v\n", VarInt)
	fmt.Printf("Type Int : %T\n", VarInt)
	fmt.Printf("Value String : %v\n", varString)
	fmt.Printf("Type String : %T\n", varString)
	fmt.Printf("Value Float : %v\n", varFloat)

}

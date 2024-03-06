package main

import "fmt"

type Supe struct { // Struct name capitalized to make it public
	name     string // name of the superhero
	strength int    // strength of the superhero
}

func main() {
	var shazam Supe              // create a new superhero
	shazam.name = "Shazam"       // set the name of the superhero
	shazam.strength = 100        // set the strength of the superhero
	fmt.Println(shazam)          // print the superhero
	fmt.Println(shazam.name)     // print the name of the superhero
	fmt.Println(shazam.strength) // print the strength of the superhero

	son := shazam          // create another superhero
	son.name = "shazamson" // create a new superhero and set it to the value of shazam
	son.strength = 1000    // set the strength of the superhero
	fmt.Println(son)       // print the superhero
	fmt.Println(son.name)  // print the name of the superhero
	fmt.Println(shazam.strength)
}

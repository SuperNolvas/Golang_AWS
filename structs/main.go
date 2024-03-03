package main

import "fmt"

type Supe struct { // Struct name capitalized to make it public
	name     string // name of the superhero
	strength int    // strength of the superhero
}

func main() {
	var shazam supe        // create a new superhero
	shazam.name = "Shazam" // set the name of the superhero
	shazam.strength = 100  // set the strength of the superhero
	fmt.Println(shazam)    // print the superhero
}

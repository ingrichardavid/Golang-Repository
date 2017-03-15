package main

import (
	"fmt"
	"math"
)

const pi float32 = math.Pi

func main() {

	//Pi is exported name and pi don't is exported name. This is exported the package math.
	fmt.Println("This is the number PI exported of the package math:", pi)
}

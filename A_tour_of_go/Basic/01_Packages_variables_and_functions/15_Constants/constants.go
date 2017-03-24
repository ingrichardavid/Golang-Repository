package main

import (
	"fmt"
	"strconv"
)

const euler float32 = 2.719

func main() {
	const greeting string = "¡Hola mundo!"
	fmt.Println(greeting + " El número euler es: " + strconv.FormatFloat(float64(euler), 'f', 3, 64))
}

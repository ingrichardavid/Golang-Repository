package main

import "fmt"

func main() {
	defer fmt.Println("Último en ejecutarse")

	fmt.Println("1")

	defer fmt.Println("Penúltimo en ejecutarse")

	fmt.Println("2")

	defer fmt.Println("Primero en ejecutarse")

	fmt.Println("3")
}

package main

import "fmt"

var (
	name            = "Luis"
	age             = 20
	salary          = 2345710.3567
	status          = true
	bonusPercentage = 0.03
	bonus           = 0.5 + 0.2i
)

func main() {
	fmt.Printf("\nname es de tipo: %T\n", name)
	fmt.Printf("age es de tipo: %T\n", age)
	fmt.Printf("salary es de tipo: %T\n", salary)
	fmt.Printf("status es de tipo: %T\n", status)
	fmt.Printf("bonusPercentage es de tipo: %T\n", bonusPercentage)
	fmt.Printf("bonus es de tipo: %T\n", bonus)
	fmt.Println()
}

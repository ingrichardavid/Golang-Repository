package main

import "fmt"

var (
	name   string
	age    int
	salary float64
	status bool
)

func main() {
	fmt.Println("\nVariables sin valor inicial: ")
	fmt.Println("\nNombre: ", name)
	fmt.Println("Edad: ", age)
	fmt.Println("Salario: ", salary)
	fmt.Println("Estatus: ", status)
	fmt.Println()

	name = "Richard David"
	age = 30
	salary = 100
	status = true

	fmt.Println("Varaibles inicializadas: ")
	fmt.Println("\nNombre: ", name)
	fmt.Println("Edad: ", age)
	fmt.Println("Salario: ", salary)
	fmt.Println("Estatus: ", status)
	fmt.Println()
}

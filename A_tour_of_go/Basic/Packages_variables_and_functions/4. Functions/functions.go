package main

import (
	"fmt"
)

func user(name string, lastName string, age string) string {
	return "Nombre: " + name + ", Apellido: " + lastName + ", Edad: " + age
}

func main() {
	fmt.Println(user("Richard", "David", "10/06/1990"))
}

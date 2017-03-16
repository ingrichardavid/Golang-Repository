package main

import (
	"fmt"
)

func user(name string, lastName string, birthday string) string {
	return "Nombre: " + name + ", Apellido: " + lastName + ", Edad: " + birthday
}

func main() {
	fmt.Println(user("Richard", "David", "10/06/1990"))
}

package main

import "fmt"

func person(name, lastName string, age int) (int, string) {
	return age, name + " " + lastName
}

func main() {
	fmt.Println(person("Miguel", "Benavides", 30))
}

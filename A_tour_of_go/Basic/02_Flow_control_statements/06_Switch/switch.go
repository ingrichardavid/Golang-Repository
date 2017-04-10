package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	blue = iota
	red
	green
	pink
	gray
)

var reader = bufio.NewReader(os.Stdin)

func getRoom() int {
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		input = strings.Replace(input, "\n", "", -1)

		room, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(err)
		} else {
			return room
		}
	}

	return 0
}

func selectedRoom(room int) {
	fmt.Println()

	switch room {
	case blue:
		fmt.Println("Seleccionó la habitación Azul.")
	case red:
		fmt.Println("Seleccionó la habitación Roja.")
	case green:
		fmt.Println("Seleccionó la habitación Verde.")
	case pink:
		fmt.Println("Seleccionó la habitación Rosado.")
	case gray:
		fmt.Println("Seleccionó la habitación Gris.")
	default:
		fmt.Println("El número no esta asociado a ninguna habitación.")
	}

	fmt.Println()
}

func main() {
	fmt.Println()
	fmt.Println("Habitaciones: ")
	fmt.Println()
	fmt.Println("\t1. Habitación Azul")
	fmt.Println("\t2. Habitación Roja")
	fmt.Println("\t3. Habitación Verde")
	fmt.Println("\t4. Habitación Rosa")
	fmt.Println("\t5. Habitación Gris")
	fmt.Println()
	fmt.Print("Seleccione su habitación: ")
	selectedRoom(getRoom())
}

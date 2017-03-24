package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func result(value int) (x, y float64) {
	x = math.Sqrt(float64(value))
	y = x + (x * 2)
	return
}

func readData() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduzca un número entero positivo: ")
	numero, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		numero = strings.Replace(numero, "\n", "", -1)
		i, err := strconv.Atoi(numero)
		if err != nil {
			fmt.Println(err)
		} else {
			if i < 0 {
				fmt.Println("Por favor, introduzca un número entero positivo.")
				readData()
			} else {
				fmt.Println(result(i))
			}
		}
	}
}

func main() {
	readData()
}

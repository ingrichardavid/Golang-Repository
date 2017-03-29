package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	limit  int
)

func requestLimit() {
	fmt.Print("\nIntroduzca el número de iteraciones: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		limit = 0
	} else {
		input = strings.Replace(input, "\n", "", -1)
		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(err)
			limit = 0
		} else {
			limit = number
		}
	}
}

func main() {
	requestLimit()

	if limit <= 0 {
		fmt.Println("\nPor favor, introduzca un número mayor a 0.")
		requestLimit()
	} else {
		fmt.Println("\nNúmero de iteraciones:" + strconv.Itoa(limit))
		fmt.Println()
		for i := 0; i < limit; i++ {
			fmt.Println("Iteración número: " + strconv.Itoa(i+1))
			fmt.Println()
		}
	}
}

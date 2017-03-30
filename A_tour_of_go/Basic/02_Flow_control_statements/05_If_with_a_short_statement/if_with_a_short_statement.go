package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Print("\nIntroduzca un texto: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		input = strings.Replace(input, "\n", "", -1)

		if numberOfCharacters := len([]rune(input)); numberOfCharacters <= 0 {
			fmt.Println("\nDebe introducir un texto.")
		} else {
			fmt.Printf("\nNúmero de caracteres: %v", numberOfCharacters)
			fmt.Println()
			fmt.Println()
		}
	}
}

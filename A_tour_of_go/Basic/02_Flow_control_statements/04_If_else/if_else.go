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
		fmt.Println("\nTexto: " + input)
		fmt.Println()
	}
}

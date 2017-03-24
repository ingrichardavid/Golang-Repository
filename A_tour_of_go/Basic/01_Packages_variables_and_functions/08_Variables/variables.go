package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	cOption = iota
	javaOption
	phpOption
	pythonOption
	objetiveCOption
	swiftOption
)

var c, java, php, python, objetiveC, swift bool

func selectOption(option int) {
	switch option {
	case cOption:
		c = true
	case javaOption:
		java = true
	case phpOption:
		php = true
	case pythonOption:
		python = true
	case objetiveCOption:
		objetiveC = true
	case swiftOption:
		swift = true
	}

	fmt.Println("\n Lenguaje de Programación seleccionado: ")
	fmt.Println("\n\t0. C: ", c)
	fmt.Println("\t1. JAVA: ", java)
	fmt.Println("\t2. Php: ", php)
	fmt.Println("\t3. Python: ", python)
	fmt.Println("\t4. Objetive-C: ", objetiveC)
	fmt.Println("\t5. Swift: ", swift)
	fmt.Println()
}

func programmingLenguage() {
	fmt.Println("\n Lenguajes de Programación: ")
	fmt.Println("\n\t0. C.")
	fmt.Println("\t1. JAVA.")
	fmt.Println("\t2. Php.")
	fmt.Println("\t3. Python.")
	fmt.Println("\t4. Objetive-C.")
	fmt.Println("\t5. Swift.")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n Seleccione su Lenguaje de Programación: ")

	option, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	} else {
		option = strings.Replace(option, "\n", "", -1)
		i, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println()
			fmt.Println("\tERROR: ", err)
			fmt.Println("\n Por favor, intente nuevamente.")
			fmt.Println()
		} else {
			if i < 0 || i > 5 {
				fmt.Println("\n Por favor, seleccione una opción válida.")
				programmingLenguage()
			} else {
				selectOption(i)
			}
		}
	}
}

func main() {
	programmingLenguage()
}

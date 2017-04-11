package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	menuTitle       string = "Menú: "
	menuOptionOne   string = "1. Sumar."
	menuOptionTwo   string = "2. Restar."
	menuOptionThree string = "3. Multiplicar."
	menuOptionFour  string = "4. Dividir."

	optionAdd      int = 1
	optionSubtract int = 2
	optionMultiply int = 3
	optionDivide   int = 4

	messageResult            string = "Resultado: "
	messageSelectAnOption    string = "Seleccione una operación: "
	messageSelectValidOption string = "Seleccione una opción válida!"
	messageEnterNumber       string = "Ingrese un número: "
	messageError             string = "Ha ocurrido un error: "
)

var reader = bufio.NewReader(os.Stdin)

func getNumber() (float32, error) {
	fmt.Println()
	fmt.Print(messageEnterNumber)

	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, errors.New(messageError + err.Error())
	}

	value, err := strconv.ParseFloat(strings.Replace(input, "\n", "", -1), 32)

	if err != nil {
		return 0, errors.New(messageError + err.Error())
	}

	return float32(value), nil
}

func operation(option int, numberOne float32, numberTwo float32) (float32, error) {
	switch {
	case option == optionAdd:
		return numberOne + numberTwo, nil
	case option == optionSubtract:
		return numberOne - numberTwo, nil
	case option == optionMultiply:
		return numberOne * numberTwo, nil
	case option == optionDivide:
		return numberOne / numberTwo, nil
	default:
		return 0, errors.New(messageSelectValidOption)
	}
}

func printMessageln(message string) {
	fmt.Println()
	fmt.Println(message)
	fmt.Println()
}

func main() {
	numberOne, errorNumberOne := getNumber()

	if errorNumberOne != nil {
		printMessageln(errorNumberOne.Error())
		return
	}

	numberTwo, errorNumberTwo := getNumber()

	if errorNumberTwo != nil {
		printMessageln(errorNumberTwo.Error())
		return
	}

	fmt.Println()
	fmt.Println(menuTitle)
	fmt.Println()
	fmt.Println("\t" + menuOptionOne)
	fmt.Println("\t" + menuOptionTwo)
	fmt.Println("\t" + menuOptionThree)
	fmt.Println("\t" + menuOptionFour)
	fmt.Println()
	fmt.Print(messageSelectAnOption)

	input, err := reader.ReadString('\n')

	if err != nil {
		printMessageln(messageError + err.Error())
		return
	}

	option, err := strconv.Atoi(strings.Replace(input, "\n", "", -1))

	if err != nil {
		printMessageln(messageError + err.Error())
		return
	}

	result, err := operation(option, numberOne, numberTwo)

	if err != nil {
		printMessageln(messageError + err.Error())
		return
	}

	fmt.Println()
	fmt.Print("\t"+messageResult, result)
	fmt.Println()
	fmt.Println()

}

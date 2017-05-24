package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	variableOne   int    = 1
	variableTwo   int    = 2
	variableThree string = "Richard"
	variableFour  string = "David"
)

func main() {
	println()
	println("******    VARIABLES DECLARADAS    ******")
	println()
	println("Variable 1: " + strconv.Itoa(variableOne))
	println("Variable 2: " + strconv.Itoa(variableTwo))
	println("Variable 3: " + variableThree)
	println("Variable 4: " + variableFour)
	println()

	pOne := &variableOne
	pTwo := &variableTwo
	pThree := &variableThree
	pFour := &variableFour

	println("******    PUNTEROS APUNTANDO A VARIABLES    ******")
	println()
	println("Puntero 1: " + strconv.Itoa(*pOne))
	println("Puntero 2: " + strconv.Itoa(*pTwo))
	println("Puntero 3: " + *pThree)
	println("Puntero 4: " + *pFour)
	println()

	*pOne = int(time.Now().Local().Month())
	*pTwo = int(time.Now().Local().Year())
	*pThree = "******    FECHA ACTUAL    ******"
	*pFour = "La fecha actual es: "

	println("******    MODIFICANDO PUNTEROS    ******")
	println()
	println("Puntero 1: " + strconv.Itoa(variableOne))
	println("Puntero 2: " + strconv.Itoa(variableTwo))
	println("Puntero 3: " + variableThree)
	println("Puntero 4: " + variableFour)
	println()

	println("******    SALIDA USANDO LAS VARIABLES    ******")
	println()
	println(variableThree)
	println()
	fmt.Printf("%s  ->>> Mes: %d, Año %d", variableFour, variableOne, variableTwo)
	println()
	println()
}

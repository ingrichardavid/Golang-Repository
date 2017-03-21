package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	bonus              float32 = 0.03
	bonusForDaysWorked int     = 20
)

var reader = bufio.NewReader(os.Stdin)

func caculateSalary(dayliSalary float32, daysWorked int) float32 {
	if daysWorked >= bonusForDaysWorked {
		return dayliSalary*float32(daysWorked) + ((dayliSalary * float32(daysWorked)) * bonus)
	}
	return dayliSalary * float32(daysWorked)
}

func getEmployeeSalary() {
	fmt.Print("\nIntroduzca el nombre del empleado: ")

	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		getEmployeeSalary()
	} else {
		name = strings.Replace(name, "\n", "", -1)
		fmt.Println("\nAl empleado " + name + " le corresponde un salario mesual de " + strconv.FormatFloat(float64(caculateSalary(getSalary(), getDaysWorked())), 'f', 3, 64))
		fmt.Println()
	}
}

func getSalary() float32 {
	fmt.Print("\nIntroduzca el salario diario: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return getSalary()
	}

	input = strings.Replace(input, "\n", "", -1)
	salary, err := strconv.ParseFloat(input, 32)

	if err != nil {
		fmt.Println(err)
		return getSalary()
	}

	if salary <= 0 {
		fmt.Println("\nEl salario diario debe ser mayor a 0.")
		return getSalary()
	}

	return float32(salary)
}

func getDaysWorked() int {
	fmt.Print("\nIntroduzca el número de días trabajados: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return getDaysWorked()
	}

	input = strings.Replace(input, "\n", "", -1)
	days, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
		return getDaysWorked()
	}

	if days <= 0 {
		fmt.Println("\nPor favor, introduzca el número de días.")
		return getDaysWorked()
	}

	return days
}

func main() {
	getEmployeeSalary()
}

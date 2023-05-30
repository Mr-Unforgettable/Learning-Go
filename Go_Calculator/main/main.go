package main

import (
	"bufio"
	"fmt"
	"operations"
	"os"
	"strconv"
	"strings"
)

var operationMap = map[string]func(float64, float64) float64{
	"+": operations.Addition,
	"-": operations.Subtraction,
	"*": operations.Multiplication,
	"/": operations.Division,
}

func main() {
	fmt.Println("===============================")
	fmt.Println("======== Go Calculator ========")
	fmt.Println("===============================")
	fmt.Println("Press 1: Addition (+)")
	fmt.Println("Press 2: Subtraction (-)")
	fmt.Println("Press 3: Multiplication (*)")
	fmt.Println("Press 4: Division (/)")
	fmt.Println("Press 0: Exit")
	fmt.Println("===============================")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Select the operation mode: ")
		modeInput, _ := reader.ReadString('\n')
		modeInput = strings.TrimSpace(modeInput)
		key, err := strconv.Atoi(modeInput)
		if err != nil {
			fmt.Println("Invalid option. Please try again.")
			continue
		}

		if key == 0 {
			fmt.Println("Exiting the program...")
			break
		}

		expression := getExpression(key)
		if expression == "" {
			fmt.Println("Invalid option. Please try again.")
			fmt.Println("=================================")
			continue
		}

		fmt.Print("Enter the number1: ")
		num1Input, _ := reader.ReadString('\n')
		num1Input = strings.TrimSpace(num1Input)
		number1, err := strconv.ParseFloat(num1Input, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			fmt.Println("=================================")
			continue
		}

		fmt.Print("Enter the number2: ")
		num2Input, _ := reader.ReadString('\n')
		num2Input = strings.TrimSpace(num2Input)
		number2, err := strconv.ParseFloat(num2Input, 64)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			fmt.Println("=================================")
			continue
		}

		if expression == "/" && number2 == 0 {
			fmt.Println("Division by zero is not allowed. Please try again.")
			fmt.Println("=================================")
			continue
		}

		result := operationMap[expression](number1, number2)
		fmt.Printf("The result is: %.2f\n\n", result)
		fmt.Println("=================================")
	}

	fmt.Println("Program exited.")
}

func getExpression(key int) string {
	switch key {
	case 1:
		return "+"
	case 2:
		return "-"
	case 3:
		return "*"
	case 4:
		return "/"
	default:
		return ""
	}
}

package operations

var operationMap = map[string]func(float64, float64) float64{
	"+": Addition,
	"-": Subtraction,
	"*": Multiplication,
	"/": Division,
}

// Calculation performs the arithmetic calculation based on the expression.
func Calculation(expression string, num1, num2 float64) float64 {
	if operation, ok := operationMap[expression]; ok {
		return operation(num1, num2)
	}
	return 0
}

// Addition adds two numbers.
func Addition(num1, num2 float64) float64 {
	return num1 + num2
}

// Subtraction subtracts two numbers.
func Subtraction(num1, num2 float64) float64 {
	return num1 - num2
}

// Multiplication multiplies two numbers.
func Multiplication(num1, num2 float64) float64 {
	return num1 * num2
}

// Division divides two numbers.
func Division(num1, num2 float64) float64 {
	if num2 == 0 {
		return 0
	}
	return num1 / num2
}

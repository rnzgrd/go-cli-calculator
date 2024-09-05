package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to perform the calculation
func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("invalid operator")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Simple CLI Calculator")
	fmt.Println("---------------------")

	for {
		fmt.Print("Enter first number: ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		fmt.Print("Enter operator (+, -, *, /): ")
		operator, _ := reader.ReadString('\n')
		operator = strings.TrimSpace(operator)

		fmt.Print("Enter second number: ")
		input2, _ := reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)
		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		result, err := calculate(num1, operator, num2)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)

		fmt.Print("Do you want to perform another calculation? (yes/no): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		if strings.ToLower(choice) != "yes" {
			fmt.Println("Exiting calculator. Goodbye!")
			break
		}
	}
}

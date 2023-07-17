package main

import "fmt"

func main() {
	// define the variables we want to add
	var number1, number2, number3 int

	// initializing the variables
	number1 = 99
	number2 = 5000

	// adding the numbers
	number3 = number1 + number2

	// printing the results
	fmt.Println("The addition of ", number1, " and ", number2, " is \n ", number3, "\n(Addition of two integers within the function)")
}

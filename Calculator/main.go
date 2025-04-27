package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	summe := x + y
	return summe
}

func subtract(x, y int) int {
	summe := x - y
	return summe
}

func multiply(x, y int) int {
	summe := x * y
	return summe
}

func divide(x, y float32) float32 {
	summe := x / y
	return summe
}

func calculator(option int) {
	switch option {
	case 1:
		var x, y int
		fmt.Println("Enter your First and Second Value")
		fmt.Scanln(&x, &y)
		summe := add(x, y)
		fmt.Println(summe)
	case 2:
		var x, y int
		fmt.Println("Enter your First and Second Value")
		fmt.Scanln(&x, &y)
		summe := subtract(x, y)
		fmt.Println(summe)
	case 3:
		var x, y int
		fmt.Println("Enter your First and Second Value")
		fmt.Scanln(&x, &y)
		summe := multiply(x, y)
		fmt.Println(summe)
	case 4:
		var x, y float32
		fmt.Println("Enter your First and Second Value")
		fmt.Scanln(&x, &y)
		summe := divide(x, y)
		fmt.Println(summe)
	default:
		errror := errors.New("not Working")
		fmt.Println(errror)
	}

}

func main() {
	var choice int
	fmt.Println("Choose a Option", "\n", "1: Add", "\n", "2: Subtract", "\n", "3: Multiply", "\n", "4: Divide")
	fmt.Scanln(&choice)
	calculator(choice)
}

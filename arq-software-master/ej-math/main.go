package main

import (
	"fmt"
	//"github.com/emikohmann/arq-software/ej-math/div"
	//"github.com/emikohmann/arq-software/ej-math/sum"
)

func main() {
	// defers the execution of a function to the end
	defer func() {
		fmt.Println("Bye!")
	}()

	// inputs
	a, b := float64(20), float64(10)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// division
	resultDiv, err := div.Division(a, b)
	if err != nil {
		fmt.Println("Error in division: ", err.Error())
		return
	}
	fmt.Println("Div: ", resultDiv)

	// sum
	resultSum := sum.Sum(a, b)
	fmt.Println("Sum: ", resultSum)
}

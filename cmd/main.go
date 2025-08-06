package main

import (
	"fmt"

	"github.com/lucasarieta/poa/solution"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("fatorial:", solution.Fatorial(10))
	fmt.Println("somatorio:", solution.Somatorio(5))
	fmt.Println("fibonacci:", solution.Fibonacci(10))
	fmt.Println("somatorio entre nums:", solution.SomatorioEntreNums(1, 5))
	fmt.Println("palindrome:", solution.Palindrome("radar"))
	fmt.Println("palindrome:", solution.Palindrome("michaelmora"))
}

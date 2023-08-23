package recursion

import "fmt"

func Fibonacci(n int) int {
	if n <= 0 && n%1 != 0 {
		fmt.Println("Valor menor que zero não permitido")
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

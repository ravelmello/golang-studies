package main

import "fmt"

func fibo(n int) int {
	if n <= 0 && n%1 != 0 {
		fmt.Println("Valor menor que zero não permitido")
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func main() {
	fmt.Println(fibo(7))
}

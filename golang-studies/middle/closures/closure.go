package main

import "fmt"

func count() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	counter := count()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

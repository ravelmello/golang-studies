package main

import "fmt"

func main() {

	a, b, c := multipleValues()

	fmt.Println("Value a:", b)
	fmt.Println("Value b:", a)

	_, _, c = multipleValues()

	fmt.Println("Value c:", c)

}

func multipleValues() (int, int, int) {
	return 2, 4, 6
}

package main

import (
	"closures"
	"fmt"
	"recursion"
	"slicesTests"
)

func main() {
	counter := closures.Count()

	fmt.Println("Using a factorial ", recursion.Factorial(5))
	fmt.Println("Using a closure ", counter())
	fmt.Println("Using a closure ", counter())
	fmt.Println("Using a closure ", counter())
	fmt.Println("CAll the slices => ", slicesTests.GiveMeASlice())

}

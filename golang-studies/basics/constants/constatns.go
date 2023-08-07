package main

import (
	"fmt"
	"math"
)

const s string = "CONSTANT"

func main() {
  fmt.Println(s)

	const n = 400000000
	const pi = 3.1
	const d = 3e20/n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sincos(45))
}
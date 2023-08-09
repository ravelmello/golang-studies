package main

import (
	"fmt"
	"strconv"
)

func main() {

	m := make(map[string]int)

	i := 0
	for i <= 10 {

		s := strconv.Itoa(i)
		m["key"+s] = i
		i++
	}

	fmt.Println(m)
}

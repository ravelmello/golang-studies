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

	m2 := make(map[string]string)
	m2["Name"] = "Ravel da Costa Melo"

	fmt.Println(m2)

	m3 := map[string]string{"item": "apple"}

	fmt.Println("map: ", m3)

}

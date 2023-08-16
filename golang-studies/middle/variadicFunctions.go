package main

import "fmt"

func soma(numbers ...int) {
	total := 0

	for _, num := range numbers {
		total += num
	}

	fmt.Println("Total da soma: ", total)
}

func main() {
	soma(1, 2, 3, 4)

	s := make([]int, 5)
	s = append(s, 1, 2, 3, 4)

	soma(s...)

	m := map[string]int{"a": 10, "b": 20}

	mapSum := make([]int, len(m))

	for _, num := range m {
		mapSum = append(mapSum, num)
	}

	soma(mapSum...)

}

package main

import "fmt"

func main() {

	numbers := []int{2, 3, 4}
	fruits := map[string]string{"a": "apple", "b": "banana", "c": "coconut"}

	multiplication := 0

	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			multiplication = numbers[i]
		} else {
			multiplication *= numbers[i]
		}
	}

	fmt.Println("Multiplication = ", multiplication)
	find(numbers)
	findInMap(fruits)

}

func findInMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("%s -> %s\n", key, value)
	}
}

func find(numbers []int) {
	for i, num := range numbers {
		if num == 3 {
			fmt.Println("Index = ", i)
		}
	}
}

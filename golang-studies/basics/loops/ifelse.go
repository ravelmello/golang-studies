package main

import "fmt"

func main() {
 
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	validateIfIsNegative(3,6)
}


func validateIfIsNegative(a int, b int)  {
	sub := a - b
	
	if sub < 0 {
    fmt.Println("The substraction of ", a, "-", b, " is negative")
  } else {
		fmt.Println("The substraction of ", a, "-", b, " is negative")
	}
}
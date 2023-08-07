package main

import "fmt"

func main() {
  var a[5] int

  fmt.Println("Empty: ",a)

	a[3] = 100

	fmt.Println("Filled: ",a)
	fmt.Println("Length: ",len(a))

	var twoDimensional[3][3] int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
      twoDimensional[i][j] = i + j
    }
	}


	fmt.Println(twoDimensional)

	fmt.Println("Get pairs: ", twoDimensional[0][1])

	b:=[5]int{1,2,3,4,5}
	fmt.Println("Non empty: ",b)
}
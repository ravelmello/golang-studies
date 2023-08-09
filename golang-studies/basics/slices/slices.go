package main

import "fmt"

func main() {

	s := make([]string, 3)

	fmt.Println("Empty", s)

	s = append(s, "d")
	s = append(s, "e")
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)

	ctrlv := make([]string, len(s))
	copy(ctrlv, s)

	fmt.Println(ctrlv)

	l := s[2:3]
	fmt.Println("Slice 1: ", l)

	l = s[2:]
	fmt.Println("Slice 2: ", l)

	l = s[:3]
	fmt.Println("Slice 3: ", l)

	fmt.Println("Length: ", len(ctrlv))

	sliceApp := make([]int, 10)

	for i := 0; i < len(sliceApp); i++ {
		sliceApp[i] = i * 2
	}

	twoD := make([][]int, 4)

	for i := 0; i < 4; i++ {
		ll := i + 1
		twoD[i] = make([]int, ll)
		fmt.Println("Lenght[i]", len(twoD[i]))
		for j := 0; j < ll; j++ {
			twoD[i][j] = i + j
		}
	}

	extendSlice := make([]int, len(sliceApp))

	copy(extendSlice, sliceApp)

	extendSlice = append(extendSlice, 10)

	fmt.Println("Appended and copy ", extendSlice)
	fmt.Println("Two Dimenensions", twoD)
}

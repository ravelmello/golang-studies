package loops

import "fmt"

func main() {
  i:=0
	for i <=10  {
		fmt.Println(i)
		i = i + 1
	}

	for j:=0; j<10; j++ {
		fmt.Println(j)
	}

	nomes := []string{"Ravel", "Karol", "Miguel", "Alice"}

	interable(nomes)

}


func interable(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}
 


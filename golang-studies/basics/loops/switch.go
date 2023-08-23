package loops

import (
	"fmt"
	"time"
)

func Switch() {
	i := 2

	fmt.Println(i, "por extenso é ")
	switch i {
	case 1:
		fmt.Println("um")
	case 2:
		fmt.Println("dois")
  case 3:	
    fmt.Println("tres")
	}

	t := time.Now()

	if t.Hour() < 12 && t.Hour() > 0 {
		fmt.Println("Bom dia")
	} else {
		fmt.Println("Boa tarde")
	}
}
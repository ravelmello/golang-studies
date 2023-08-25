package main

import (
	"closures"
	"fmt"
	"recursion"
	"slicesTests"
	"structs"
)

func main() {
	counter := closures.Count()

	fmt.Println("Using a factorial ", recursion.Factorial(5))
	fmt.Println("Using a closure ", counter())
	fmt.Println("Using a closure ", counter())
	fmt.Println("Using a closure ", counter())
	fmt.Println("CAll the slices => ", slicesTests.GiveMeASlice())

	pf := structs.PhisicalPerson{
		Person: structs.Person{Name: "Ravel",
			Status: true,
			Age:    31},
		Surname: " da Costa Melo",
		CPF:     "000.000.000-00",
	}

	pj := structs.JuridicPerson{
		Person:     structs.Person{Name: "Ravel INC", Age: 2, Status: true},
		TaxAddress: "Avenue 2st",
		CNPJ:       "00.000.000./0000-00",
	}

	fmt.Println("Make a person using structs => ", pf)
	fmt.Println("Get document from interface impl => ", structs.Show(pf))

	fmt.Println("Make a person using structs => ", pj)
	fmt.Println("Get document from interface impl => ", structs.Show(pj))

}

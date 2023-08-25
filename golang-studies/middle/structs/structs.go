package structs

import (
	"fmt"
	"time"
)

type Pessoa struct {
	Name      string
	Age       int8
	isLive    bool
	BirthDate time.Time
	cpf       string
}

func makePerson() Pessoa {
	p := Pessoa{
		Name:      "Ravel Melo",
		isLive:    true,
		Age:       31,
		BirthDate: time.Date(1992, 01, 04, 00, 00, 00, 00, time.UTC),
		cpf:       "000.000.000-00",
	}

	return p

}

func GiveMeAPerson() Pessoa {
	return makePerson()
}


func (p Pessoa) String() string {
	return fmt.Sprintf("Name: %s  Age: %d isAlive: %t", p.Name, p.Age, p.isLive)
}

package structs

import "fmt"

type Person struct {
	Name   string
	Status bool
	Age    int8
}

type JuridicPerson struct {
	Person
	CNPJ       string
	TaxAddress string
}

type PhisicalPerson struct {
	Person
	Surname string
	CPF     string
}

func (pp PhisicalPerson) Doc() string {
	return fmt.Sprintf("Document CPF: %s", pp.CPF)
}

func (jp JuridicPerson) Doc() string {
	return fmt.Sprintf("Document CNPJ: %s", jp.CNPJ)
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s  Age: %d isAlive: %t", p.Name, p.Age, p.Status)
}

func Show(d Document) string {
	return d.Doc()
}

package main

import "fmt"

type Teste struct {
	ID int
}

func (t Teste) PrintName() string {
	return fmt.Sprintf("ID: %d", t.ID)
}

func print(t Teste) {
	fmt.Println(t.PrintName())
}

func main() {

	var teste3 Teste

	print(teste3)

}

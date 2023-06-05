package main

import "fmt"

/*

Interfaces
São estruturas que definem contratos que outras estruturas ou tipos
devem seguir, implementando os métodos da mesma.

Interfaces definem um conjunto de métodos que um tipo de dado deve implementar
para ser considerado membro da interface.

*/

type Animal interface {
	Speak() string
}

// Criar uma struct que implementa o método do contrato

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {

	var cachorro Animal = Dog{Name: "Betovem"}
	fmt.Println(cachorro)
	fmt.Println(cachorro.Speak())

}

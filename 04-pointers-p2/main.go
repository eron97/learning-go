package main

import "fmt"

func doubleValue(value *int) {
	*value *= 2
}

func main() {

	// Ponteiros para tipos de dados básicos!
	var intValue int = 10
	var intPointer *int // criando uma variavel ponteiro do tipo de um inteiro
	intPointer = &intValue

	fmt.Printf("intValue está alocado em %p", &intValue)
	fmt.Println()
	fmt.Printf("intPointer está alocado em %p", intPointer)
	fmt.Println()
	// Passagem de ponteiro como parâmetro de função

	fmt.Println("Valor de intValue antes:", intValue)

	doubleValue(&intValue)
	fmt.Println("Depois:", intValue)

}

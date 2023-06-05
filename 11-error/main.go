package main

import "fmt"

/*

A função Errorf é usada para criar uma mensagem de erro formatada.
Ela retorna um valor do tipo error, que é uma interface comummente
usada para representar erros em Go.

*/

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divisão por zero")
	}
	return a / b, nil
}

func main() {

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", result)
	}

}

package main

import "fmt"

/*

O método Scan é usado para ler e armazenar valores de entrada do usuário.
Ele lê os dados fornecidos pelo usuário a partir da entrada padrão
(normalmente o teclado) e armazena esses valores nas variáveis fornecidas
como argumentos.

*/

func main() {

	var name string
	var age int

	fmt.Println("Digite seu nome e idade:")
	fmt.Scan(&name, &age)

	fmt.Printf("Seu nome é %s e você tem %d anos de idade", name, age)

}

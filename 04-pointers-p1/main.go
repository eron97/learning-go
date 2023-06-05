package main

import "fmt"

/*

Em Go, um ponteiro é uma variável que armazena o endereço
de memória de outra variável.

Ponteiros permitem manipular diretamente objetos na memória e
melhorar a eficiência de um programa.

Quando as duas variáveis compartilham o mesmo valor e, ao alterar
uma das variáveis, a outra também é alterada, isso é chamado
de "passagem por referência". As duas variáveis apontam para o
mesmo endereço de memória.

Em relação ao armazenamento e controle de ponteiros, o Go possui um coletor
de lixo (garbage collector) embutido que gerencia a alocação e a
liberação de memória.

Em relação ao armazenamento e controle de ponteiros, o Go possui um
coletor de lixo (garbage collector) embutido que gerencia a alocação e a
liberação de memória.

*/

func main() {
	// Passagem por valor -> num1 não é alterado por num2
	num1 := 10
	num2 := num1

	fmt.Printf("num1: %d, num2: %d\n", num1, num2)
	fmt.Printf("num1: %d, endereço de memória: %p\n", num1, &num1)
	fmt.Printf("num2: %d, endereço de memória: %p\n", num2, &num2)

	num2 = 30

	fmt.Printf("num1: %d, num2: %d\n", num1, num2)
	fmt.Printf("num1: %d, endereço de memória: %p\n", num1, &num1)
	fmt.Printf("num2: %d, endereço de memória: %p\n", num2, &num2)
	// %p verbo de formatação de impressão para exibir endereços de memória em variáveis
	// Operador & para obter o endereço de uma variável

	// Passagem por referência
	nomes := []string{"João", "Maria"}

	fmt.Println(nomes[0])
	fmt.Println(nomes[1])
	fmt.Printf("Slice nomes, endereço de memória: %p\n", &nomes)

	novosNomes := nomes
	fmt.Println(novosNomes[0])
	fmt.Println(novosNomes[1])
	fmt.Printf("Slice novosNomes, endereço de memória: %p\n", &novosNomes)

	novosNomes[0] = "Marcelina"

	fmt.Println(nomes)
	fmt.Println(novosNomes)

}

/*

OBS

Quando você atribui um array a outra variável, é feita uma cópia completa
do conteúdo do array, e não uma referência ao array original.

Os slices são estruturas de dados que fornecem uma visualização dinâmica
de um segmento de um array subjacente. Ao atribuir um slice a outra
variável, ambas as variáveis apontam para a mesma estrutura de dados
subjacente.

Caso eu tenha um array de nomes e crio um slice a partir dele e faço
uma alteração no primeiro elemento do array isso também vai refletir no
slice, porém o contrário não acontece.

O slice é sempre uma referência ao array.


*/

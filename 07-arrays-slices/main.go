package main

import "fmt"

/*

Arrays ou vetores
-> coleção ordenada de elementos podendo ser
unidimensionais ou multidimensionais

são fixos e possuem quantidade de elementos determinada em sua criação.

*/

/*

Slice
-> uma estrutura semelhante a um arry.
Possui seu tamanho variável podendo crescer ou diminuir dinamicamente.
(É construído sobre um array subjacente)
A função Make é usada para criar um slice a partir de um array.

*/

func main() {

	var nomes [2]string
	nomes[0] = "João"
	nomes[1] = "Maria"

	/* ou...

	   nomes := [2]string{"João", "Maria"}

	*/

	// vetor bidimencional

	numeros := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(numeros)

	for _, value := range numeros {
		fmt.Println(value)
	}

	fmt.Println("------------------")

	for _, linha := range numeros {
		for _, valor := range linha {
			fmt.Println(valor)
		}
	}

	// Slice

	slice := make([]string, 2)
	slice[0] = nomes[0]
	slice[1] = nomes[1]
	fmt.Println(slice)

	slice = append(slice, "Pedro")

	for index, value := range slice {
		fmt.Println(index, value)
	}

}

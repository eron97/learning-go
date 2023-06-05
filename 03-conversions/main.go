package main

// int -> int8 -> int16 -> int32 -> int64 (positivos e negativos)
// uint -> uint8 -> uint16 -> uint32 -> uint64 (valores não negativos)

/*

 Em Go, uma conversão explícita ocorre quando você deseja converter um valor de um
 tipo para outro, desde que não haja perda de dados ou truncamento envolvido.

*/

import (
	"fmt"
	"strconv"
)

func main() {

	// Conversões explícitas

	// int para float
	var numInt int = 42
	var numFloat float64 = float64(numInt)
	fmt.Printf("Conversão explícita: int to float64 - %f\n", numFloat)

	/*

		Algumas conversões de um tipo para outro requer o uso de funções
		apropriadas para fazer a conversão correta.

	*/
	var str string = "123"

	numInt, _ = strconv.Atoi(str) // string para int com uso da função strconv.Atoi()
	fmt.Printf("Conversão explícita: string to int - %d\n", numInt)

	// rune para int
	var char rune = 'A'
	var numInt2 int = int(char)
	fmt.Printf("Conversão explícita: rune to int - %d\n", numInt2)

	// int para float64
	var numInt4 int = 42
	var numFloat2 float64
	numFloat2 = float64(numInt4)
	fmt.Printf("Conversão implícita: int to float64 - %f\n", numFloat2)

	// int para bool
	var numInt7 int = 1
	var flag2 bool = numInt7 > 0
	fmt.Printf("Conversão implícita: int to bool - %t\n", flag2)

	// bool para int
	var verdadeira = true
	var numverdadeiro int

	if verdadeira == true {
		numverdadeiro = 1
	} else {
		numverdadeiro = 0
	}

	fmt.Printf("Conversão implícita: bool to int - %d\n", numverdadeiro)

}

/*

%v ou %s: Representa o valor de um argumento de forma padrão, como uma string para strings, o valor subjacente para tipos numéricos, etc.
%d ou %i: Formata um valor como um número decimal (inteiro).
%f: Formata um valor como um número de ponto flutuante.
%t: Formata um valor booleano como true ou false.
%c: Formata um valor numérico como um caractere Unicode.
%p: Formata um ponteiro como um valor hexadecimal.
%b: Formata um valor numérico como uma representação binária.
%o: Formata um valor numérico como uma representação octal.
%x ou %X: Formata um valor numérico como uma representação hexadecimal.
%e ou %E: Formata um valor numérico usando a notação científica (exponencial).
%q: Formata um valor como uma string com aspas duplas, escapando caracteres especiais.

etc

*/

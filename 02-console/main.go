package main

import (
	"fmt"
	"reflect"
)

/*

O pacote fmt em Go fornece funcionalidades para formatar e imprimir dados no console
ou em outras saídas. Ele é amplamente utilizado para exibir informações durante a
execução de um programa.

As principais funções do pacote fmt são Println, Printf e Print.

fmt.Println: A função Println é usada para imprimir um ou mais valores no console,
adicionando uma nova linha após a impressão. É útil para imprimir mensagens ou depurar
informações.

fmt.Printf: A função Printf é usada para imprimir dados formatados com base em verbos
especificadores de formato. Você pode usar para controlar a exibição de valores, como
formatação de números, alinhamento e preenchimento.

fmt.Print: A função Print é semelhante a Println, mas não adiciona uma nova linha após
a impressão. Ela é útil quando você deseja imprimir várias informações na mesma linha.
*/

func main() {

	name := "John"
	age := 30

	// fmt.Println
	fmt.Println(name)
	fmt.Println(age)

	// fmt.Printf
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)

	// fmt.print
	fmt.Print("Name: ", name, " - ")
	fmt.Print("Age: ", age, "\n")

	// Verificando o tipo da variavel com fmt.Printf
	fmt.Printf("Tipo de 'name': %T\n", name)
	fmt.Printf("Tipo de 'age': %T\n", age)

	// Outra opção é usar a função reflect.TypeOf do pacote reflect
	fmt.Print("Reflect informa o tipo de 'name' -> ", reflect.TypeOf(name))

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

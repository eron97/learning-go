package main

import "fmt"

/*

Struct
uma estrutura que agrupa valores de diferentes tipos em um único objeto.

Func
usada para declarar uma função.
Podem ser tratadas como variáveis e passadas como argumentos
para outras funções

*/

func soma(a, b int) int {
	return a + b
}

type Pessoa struct {
	Nome  string
	Idade int
}

type PessoaComMetodos struct {
	Pessoa
}

func (p PessoaComMetodos) Andar() {
	fmt.Println("Começou a se levanter e andar lentamente!")
}

// Struct Pessoa poderia ter sido criada dentro da func main()

func main() {

	// Criando uma nova objeto do tipo Pessoa
	var p1 Pessoa
	p1.Nome = "João"
	p1.Idade = 38

	fmt.Println(p1)
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	// Somente p2 possui o método Andar()
	// Adicionando um método Associado ao objeto da struct
	p2 := PessoaComMetodos{
		Pessoa: Pessoa{
			Nome:  "Maria",
			Idade: 19,
		},
	}

	fmt.Printf("%s têm %d ", p2.Nome, p2.Idade)
	p2.Andar()

	resultado := soma(2, 3)
	fmt.Println(resultado)

}

/*

DEFER
utilizado para adiar a execução de uma função até que a função atual seja concluída

func main(){

	defer fmt.Println("deferred function")
	fmt.Println("main function")
}

saída gerada ->

main function
deferred function

*/

package main

import "fmt"

/*

Declarações de Variáveis e Constantes

Em Go, a declaração de variáveis e constantes dentro e fora de uma função é permitida,
mas existem algumas diferenças de escopo e visibilidade a serem consideradas.

Quando declaradas fora da função (a nível de pacote), sua visibilidade é
considerado do tipo global e podem ser chamadas a qualquer momento.

Variáveis podem ser declaradas de três formas
- forma explicítica
- forma por inferência de tipo
- forma curta (inicializada)
Podem ou não ter valores atribuídos a elas, pois já assumem um valor padrão.

*/

var a int = 10 // declaração explicita
var b = 10     // declaração com inferência de tipo
var c int      // declaração com inferência de tipo assumindo um valor padrão zero
var variableD string = "variableD"
var variableE string // assumindo um valor padrão vazio ""
var runeF rune = 'F'
var runeG rune // assumindo um valor padrão '\x00'

/*

Constantes são úteis para valores que não devem ser alterados durante a execução do programa.
(valores matemáticos fixos, códigos de erro, configurações globais) etc

Constantes devem sempre ser inicizadas com algum valor atribuido (não existe valor padrão)

*/

const outsideInt int = 1000

// gera erro -> const outsideString string
const outsideString string = "outsideString"

func main() {

	// Atribuição de valores à variáveis somente é possível dentro de funções, exemplo:
	variableE = "variableE"
	runeG = 'G'

	h := 100 // declaração, inicialização e atribuição de valor por inferência de tipo

	fmt.Println(variableD)
	fmt.Println(variableE)
	fmt.Println(runeF)
	fmt.Println(runeG)
	fmt.Println(h)

	// gera erro (const não deve ser alterada) -> outsideInt = 2000

	fmt.Println(outsideInt)
	fmt.Println(outsideString)

	const insideFloat64 float64 = 4.44

	fmt.Println(insideFloat64)

}

/*

BOAS PRÁTICAS??

Trabalhamos apenas com constantes a nível de pacote.

Declararamos variáveis no escopo mais restrito possível, apenas
quando são necessárias em um contexto específico.

*/

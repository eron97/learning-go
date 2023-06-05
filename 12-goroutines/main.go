package main

import (
	"fmt"
	"sync"
)

/*

Goroutine é uma função que é executada de forma assíncrona e concorrente
em relação ao resto do programa.

Goroutines permitem que o programa execute várias tarefas simultaneamente
o que faz melhroar significamente o desemepnho e eficiência do programa.

*/

func main() {

	var wg sync.WaitGroup
	wg.Add(1) // Incrementa o contador do WaitGroup

	go minhaGoroutine(&wg)
	fmt.Println("Aqui está a função principal")

	// saída apenas da função main (função principal do programa)
	// minhaGoroutine foi fechada sem ser completad. O que fazer?

	/*

		Vamos usar mecanismos de sincronização, como o wait groups.
		('sync.WaitGroup), canais ('channels') ou outros métodos.

	*/

	wg.Wait() // Aguarda a conclusão da goroutine

	fmt.Println("Fim do programa")

}

func minhaGoroutine(wg *sync.WaitGroup) {
	defer wg.Done() // Indica que a goroutine concluiu
	fmt.Println("Está é uma goroutine")
}

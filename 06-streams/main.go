package main

import "fmt"

func main() {

	// if else

	x := 11

	if x > 10 {
		fmt.Println("X é maior que 10")
	} else {
		fmt.Println("X é menor que 10")
	}

	// for classico
	for i := 0; i < 11; i++ {
		fmt.Printf("%d /", i)
	}

	/*
		while com for (n existe while em Go)

		i := 0
		for i < 11 {
			fmt.Printf("%d /", i)
			i++
		}

	*/

	// while com continue e break
	y := 2

	for y < 10 {
		y++
		if y != 8 {
			continue
		} else {
			fmt.Printf("%d ", y)
		}
	}

	fmt.Println("Loop finalizado")

	// for range em slice
	nomes := [3]string{"João", "Maria", "Pedro"}

	for index, names := range nomes {
		fmt.Println(index, names)
	}

	// swhich
	option := 2

	switch option {
	case 1:
		fmt.Println("Opção 1 escolhida")
	case 2:
		fmt.Println("Opção 2 escolhida")
	default:
		fmt.Println("Opção 0 escolhida")
	}

}

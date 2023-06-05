package main

import "fmt"

func main() {

	// Uma coleção de pares chave e valor

	var myMap map[string]int

	myMap2 := map[string]int{
		"um":   1,
		"dois": 2,
		"três": 3,
	}

	fmt.Println(myMap)
	fmt.Println(myMap2)

	/*

		Para adicionar valores a um mapa em Go, você precisa primeiro inicializá-lo
		usando a função make() e, em seguida, atribuir valores a cada chave.

	*/

	myMap = make(map[string]int)

	myMap["Chave1"] = 10
	myMap["Chave1000"] = 1000

	fmt.Println(myMap)

}

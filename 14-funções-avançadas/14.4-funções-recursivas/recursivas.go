package main

import "fmt"

// unit é ou são número(ros) positivo(vos)
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções recursivas")

	// 1 1 2 3 5 8 13

	posicao := uint(12)

	// fmt.Println(fibonacci(posicao))

	for i := uint(1); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}

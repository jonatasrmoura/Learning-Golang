/*
	O Ponteiro é uma variável que vai salvar um endereço de memória de alguma coisa
*/

package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA, ELE NÃO GUARDA VALOR DENTRO DELE E SIM REFERENCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	// Aqui retorna o valor da variavel3 e o endereço de memória de dentro do meu computador (ponteiro)
	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // desreferenciação/pegando o valor que está na referencia de memória  (* com uma variavel)
}

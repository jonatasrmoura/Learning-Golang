package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	// fixar 5 index para o array1 aonde eles só podem ser uma string
	var array1 [5]string
	array1[0] = "Posição 1"
	array1[1] = "Posição2"
	fmt.Println(array1)

	array2 := [5]string{"Posição1", "Posição2", "Posição3", "Posição4", "Posição5"}
	fmt.Println(array2)

	// Vai fixar o tamanho do array de acordo com a quantidade de items que passei dentro dele.
	/* Neste exemplo não vai deixar o array com o tamanho dinâmico, ele só deixa com o tamanho
	... baseado na quantidade de valores que eu passei para ele
	*/
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// SLICES

	/*
		O Slice é uma estrutura baseado no array, só que ela têm flexibilidades á mais principalmente
		a questão do tamanho.
		O tamanho dele pode mudar de acordo com minha necessidade, ele não têm um tamanho fixo, então
		se eu quiser usar dois valores no Slice beleza e se eu quiser colocar mais 100 nesse mesmo Slice
		também não tem problema.
	*/

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	// O Slice é uma fatia de uma array, ele pega as posição que eu quero

	// Esse Slice vai ser uma fatia do meu array2 e vai devolver a posição 2 até o 3 do array2
	// Ele não vai pegar o index 4! ele ira parar quando chegar no index 4
	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Se as posições do array mudar, também ira afetar o Slice

	array2[1] = "Posição alterada"
	fmt.Println(array2)
}

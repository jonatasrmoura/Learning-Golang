package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controles")

	numero := 0

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("Igual à 15")
	} else {
		fmt.Println("Menor que 15")
	}

	/*
		IF INIT => quando você executa uma condição if e inicializa uma variável
		já nesta condição, é uma variável de ESCOPO dessa condição
	*/

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if outroNumero == 0 {
		fmt.Println("Número é igual a zero")
	} else {
		fmt.Println("Número é menor que zero")
	}
}

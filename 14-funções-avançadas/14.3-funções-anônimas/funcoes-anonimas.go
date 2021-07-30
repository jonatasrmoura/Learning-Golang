package main

import "fmt"

func main() {
	// Essa função vai ser executada assim que for lida pelo programa
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("Passando parâmetro")

	fmt.Println(retorno)
}

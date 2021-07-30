package main

import "fmt"

func recuperarExecucao() {
	// recover para resolver problema de panic, não para a execução do programa

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	// vai interromper o fluxo normal do meu programa e vai parar tudo
	// Ele vai parar a execução do programa. Não é a melhor maneira de tratar erro do meu programa
	// Ele não é do typo Erro.
	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(8, 6))
	fmt.Println("Pós execução!")
}

package main

import (
	"fmt"
	"time"
)

// Sempre usa for para fazer loop no golang, só for

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// Posso usar variável de ambiente
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := []string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// Pegando só os nomes sem os indices

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	// Pegar cada letra da palavra "PALAVRA"
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	// usando com map
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// NÃO PODE USAR RANGE EM UM STRUCT, SÓ EM MAPS

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Loop infinito
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}

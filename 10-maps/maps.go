package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuarios := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	// mostrar o sobrenome
	fmt.Println(usuarios["sobrenome"])

	usuarios2 := map[string]map[string]string{
		"nome": {
			"Primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuarios2)

	// Deletar uma chave
	delete(usuarios2, "nome")

	// Adicionar mais uma chave
	usuarios2["signo"] = map[string]string{
		"nome": "Libra",
	}

	fmt.Println(usuarios2)
}

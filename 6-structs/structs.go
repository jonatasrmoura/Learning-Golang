package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
	email string
	senha string
	endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Jonatas"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{ "Rua dos Bobos", 120 }
	fmt.Println(enderecoExemplo)

	usuario2 := usuario{"Davi", 25, "davi@gmail.com", "123", enderecoExemplo} 
	fmt.Println(usuario2)

	usuario3 := usuario{ nome: "Jorge"}
	fmt.Println(usuario3)
}

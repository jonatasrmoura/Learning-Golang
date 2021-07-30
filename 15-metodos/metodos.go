package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

/*
	Todos os usuário têm o método salvar
*/
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Jonatas Rosa Moura", 21}
	usuario1.salvar()

	usuario2 := usuario{"Bruno Rosa Moura", 15}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println("Usuário", usuario2.nome, "possui", usuario2.idade, "anos de idade")
}

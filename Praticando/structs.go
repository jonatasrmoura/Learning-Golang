package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	age       uint8
	active    bool
}

type profission struct {
	name string
	wage float32
}

func main() {
	fmt.Println("Práticando Structs")

	var u user
	var pro profission

	u.firstName = "Jonatas"
	u.lastName = "Rosa Moura"
	u.age = 21
	u.active = true

	func userAcitve() string {
		if u.active {
			return "Trabalhando"
		}
		return "Desligado"
	}

	pro.name = "Desenvolvedor Back-end de golang"
	pro.wage = 7000

	fmt.Println(
		"Meu nome é", u.firstName, u.lastName, "e tenho", u.age, "anos de idade, sou", pro.name, "e ganho", pro.wage, "-", userAcitve(),
	)

	user2 := user{"Bruno", "Rosa Moura", 15, true}
	user2Pro := profission{"Desenvolvedor back-end júnior de C#", 5000}
	fmt.Println(user2.firstName, user2.lastName, "tem", user2.age, "anos de idade e ele é", user2Pro.name, "seu salário é de", user2Pro.wage)
}

package main

import "fmt"

var n int

// função é executada antes da func main
// Eu posso ter um func init() por arquivo, não por pacote
func init() {
	fmt.Println("Executandp a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}

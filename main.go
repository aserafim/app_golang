package main

import "fmt"

func main() {
	fmt.Println("Ponto de partida")

	var teste *int

	fmt.Println(&teste)
	fmt.Println(teste)
	inteiro := 8
	teste = &inteiro
	fmt.Println(*teste)
}

package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	var resultado bool = true
	var resto int
	for i := 2; i < num; i++ {
		resto = num % i
		if resto == 0 {
			resultado = false
		}
	}
	if resultado == true {
		fmt.Print("É primo")
	} else {
		fmt.Print("Não é primo")
	}
}

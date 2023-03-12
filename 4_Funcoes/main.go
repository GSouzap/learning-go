package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// As funções podem ter mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// Em GO, funções são tipos também
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("Texto da função 1")
	result := f("Texto da função")
	fmt.Println(result)

	resultSoma, resultSub := calculosMatematicos(10, 15)
	fmt.Println(resultSoma, resultSub)

	// Usar _ quando eu não quero usar um dos valores retornados
	resultSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultSoma2)
}
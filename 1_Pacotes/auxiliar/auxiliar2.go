package auxiliar

import "fmt"

// Como escrever2 é uma função que não será referenciada fora do pacote, ela é declarada com a primeira letra lowercase
func escrever2() {
	fmt.Println("Escrevendo pela função escrever2")
}
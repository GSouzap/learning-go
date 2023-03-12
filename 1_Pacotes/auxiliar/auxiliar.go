package auxiliar

import "fmt"

// Como Escrever é uma função que será referenciada fora do pacote, ela é declarada com a primeira letra uppercase
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	// Dentro do pacote eu consigo referenciar funções que estão com a primeira letra lowercase, fora dele não
	escrever2()
}
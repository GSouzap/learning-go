package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail" // Importação de pacotes externos
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	// Para referenciar a função fora do pacote, ela necessáriamente deve ser declarada com a primeira letra uppercase
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
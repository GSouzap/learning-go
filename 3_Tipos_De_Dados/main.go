package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos basicos de dados do GO

	// Numeros inteiros (variam a quantidade de bits que o numero pode ter)
	var numero int16 = 100
	fmt.Println(numero)
	// Caso eu não defina um valor de bits, ele usará a arquitetura da maquina como base
	var numero2 int = 1000000000
	fmt.Println(numero2)
	// Int pode receber numeros negativos
	numero3 := -100
	fmt.Println(numero3)

	// Numeros inteiros sem sinal (uint: unsigned int)
	var numero4 uint = 10
	fmt.Println(numero4)

	// Alias: "apelidos" para os tipos
	var numero5 rune = 12345 // Rune = int32
	fmt.Println(numero5)
	var numero6 byte = 123 // Byte = uint8
	fmt.Println(numero6)

	// Numero flutuantes
	var numero7 float32 = 10.5
	fmt.Println(numero7)

	// String
	var str string = "Texto"
	fmt.Println(str)

	// Declarando apenas um caracter, ele é convertido para ASCC
	char := 'B'
	fmt.Println(char)

	// Valor 0: no GO, todo tipo tem um valor inicial, ou seja, quando declarada uma variável, 
	// como padrão ela vem com o valor 0 estipulado para o tipo
	var num int
	fmt.Println(num) // Retorna 0

	var texto string
	fmt.Println(texto) // Retorna uma string vazia

	// Booleano
	var booleano bool = true
	fmt.Println(booleano) // Valor 0 false

	// Erro é um tipo próprio muito usado em GO
	var erro error = errors.New("Erro interno")
	fmt.Println(erro) // Valor 0 <nil>
}
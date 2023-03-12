package main

import "fmt"

func main() {
	// Existem duas maneiras de declarar variável em GO: deixando o tipo implícito e deixando o tipo explicito
	// Primeira forma: declarando o tipo na variável (explicito)
	var variavel1 string = "variable one"
	fmt.Println(variavel1)

	// Segunda forma: declarando o tipo de forma implicita (a variável entende o tipo pelo valor que ela recebe)
	// Inferencia de tipo
	variavel2 := "variable two"
	fmt.Println(variavel2)

	// Declarndo várias variáveis ao mesmo tempo
	var (
		variavel3 string = "lalalala"
		variavel4 string = "lalalal"
	)
	fmt.Println(variavel3, variavel4)

	// Declarndo várias variáveis ao mesmo tempo com inferencia
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	// Invetendo o valor de variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
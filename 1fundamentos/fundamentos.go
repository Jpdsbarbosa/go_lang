package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func variaveis() {

	var texto string
	var numero int
	var metro float32
	var ehFeminino bool

	fmt.Println("### Variáveis ###")
	fmt.Println(texto)
	fmt.Println(numero)
	fmt.Println(metro)
	fmt.Println(ehFeminino)
	fmt.Println("### ----------- ###")
	fmt.Println("\n")
}

func operadores() {

	fmt.Println("### Operadores ###")
	num1 := 10.0
	num2 := 20.0
	result := num1 / num2
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))
	texto1 := "texto 1"
	texto2 := "texto 2"
	resultTexto := texto1 + texto2
	fmt.Println(resultTexto)
	fmt.Println(reflect.TypeOf(resultTexto))
	fmt.Println("### ----------- ###\n")

}

func constantes() {

	fmt.Println("### Constantes ###")
	const taxa = 10 // não é possível alterar o valor de taxa por ser uma constante
	// taxa = 5 -> essa operação não é possível
	fmt.Println(taxa)
	fmt.Println(reflect.TypeOf(taxa))
	fmt.Println("### ----------- ###\n")

}

func tiposDeDados() {

	fmt.Println("### Tipos de Dados ###")
	var numero1 uint8 // não pode atribuir valores negativos para variaveis do tipo "u"
	numero1 = 1
	var numero2 int8 // inteiros de 8bits vão de -127 a 127
	numero2 = 127
	// usar o int geralmente resolve os problemas
	fmt.Println(numero1)
	fmt.Println(numero2)

	/*
		int8 	= 8-bit signed integer
		int16	= 16-bit signed integer
		int32	= 32-bit signed integer
		int64	= 64-bit signed integer
		uint8	= 8-bit unsigned integer
		uint16	= 16-bit unsigned integer
		uint32	= 32-bit unsigned integer
		uint64	= 64-bit unsigned integer
		int  	= Both in and uint contain same size, either 32 or 64 bit
		uint 	= Both in and uint contain same size, einther 32 or 64 bit

		float32 = 32-bit IEEE 754 floating-point number
		float64 = 64-bit IEEE 754 floating-point number
	*/
	fmt.Println("### ----------- ###\n")

}

func conversao() {

	fmt.Println("### Conversao ###")

	// conversão de int para int8
	var numero int = 127
	var numeroInt8 int8
	numeroInt8 = int8(numero)
	fmt.Println(numeroInt8)

	// conversão de int para float
	var numero2 float32
	numero2 = float32(numero)
	fmt.Println(numero2)

	// conversão de float para int é possivel??

	var numero3 float32
	numero3 = 127.8
	var numero4 int
	numero4 = int(numero3)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println("### ----------- ###\n")

	// pacote strconv -> para converter variáveis

	fmt.Println("### Conversao de tipos ###")

	b, _ := strconv.ParseBool("true") // "_" para "tratar um erro"
	fmt.Printf("%T \n", b)            // %T para retornar o tipo
	fmt.Println(b)

	c, _ := strconv.ParseInt("-42", 10, 64) // convertendo o -42, base 10 e 64 bits para inteiro
	fmt.Printf("%T \n", b)
	fmt.Println(c)

	texto := "42.55"

	d, _ := strconv.ParseFloat(texto, 64) // convertendo a variavel texto para float64
	fmt.Printf("%T \n", b)
	fmt.Println(d)

	fmt.Println("### ----------- ###\n")
}

func main() {

	variaveis()
	operadores()
	constantes()
	tiposDeDados()
	conversao()

}

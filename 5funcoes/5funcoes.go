package main

//funções init são funções que sempre são chamadas inicializando o arquivo

import "fmt"

func init() {

	fmt.Println("Funcão init 1")
}

func init() {

	fmt.Println("Funcão init 2")
}

func main() {

	fmt.Println("Ola")

}

func init() {

	fmt.Println("Funcão init 3")
}

func init() {

	fmt.Println("Funcão init 4")
}

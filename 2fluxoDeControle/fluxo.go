package main

import (
	"fmt"
)

func salarioComBonus() {

	salario := 999.00
	var salarioMaisBonus float64

	salarioMaisBonus = salario

	if salario <= 1000 {
		salarioMaisBonus = (salarioMaisBonus + 100)
	}

	fmt.Println("Salário:", salarioMaisBonus)
}

func carroOuMoto() {

	var ehCarro bool = true
	valorDoAutomovel := 1000.00

	if ehCarro {
		valorDoAutomovel = valorDoAutomovel + 55.50
	}

	fmt.Println("Valor Final:", valorDoAutomovel)

}

func taxaCarroOuMoto() {

	var ehCarro bool = false
	valorDoAutomovel := 1000.00

	if ehCarro {
		valorDoAutomovel += 55.00
	} else {
		valorDoAutomovel += 20.50
	}

	fmt.Println("Valor Final:", valorDoAutomovel)

}

func descontoSalarial() {

	salario := 2000.00

	if salario <= 1000.00 {
		salario -= (salario * 0.08)
	} else if salario <= 1200.00 {
		salario -= (salario * 0.10)
	} else {
		salario -= (salario * 0.12)
	}

	fmt.Println("Salario Final:", salario)

}

func descontoSalarial2() {

	salario := 10000.00
	tipo := "PJ"

	if salario < 1000.00 && tipo == "CLT" {
		salario -= (salario * 0.08)
	} else if salario < 1000 && tipo == "PJ" {
		salario -= (salario * 0.05)
	} else if salario <= 1200 && tipo == "CLT" {
		salario -= (salario * 0.10)
	} else if salario <= 1200 && tipo == "PJ" {
		salario -= (salario * 0.06)
	} else if salario > 5000 && tipo == "CLT" {
		salario -= (salario * 0.15)
	} else {
		salario -= (salario * 0.08)
	}

	fmt.Println("Salario Final:", salario)
}

func testantoFor() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func testandoBreak() {

	texto := "palavra"
	fmt.Println("Quantidade:", len(texto))

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			break
		}
		fmt.Println(string(texto[i]))
	}

}

func testandoContinue() {

	texto := "palavra"
	fmt.Println("Quantidade:", len(texto))

	for i := 0; i < len(texto); i++ {
		if string(texto[i]) == "r" {
			continue
		}
		fmt.Println(string(texto[i]))
	}
}

func fazendoWhileComFor() {

	texto := "palavra"
	fmt.Println("Quantidade:", len(texto))
	tamanho := len(texto)
	i := 0

	for i < tamanho {
		if string(texto[i]) == "r" {
			break
		}
		fmt.Println(string(texto[i]))
		i++
	}
}

func forAninhados() {

	for numBase := 1; numBase <= 10; numBase++ {

		for multiplicado := 1; multiplicado <= 10; multiplicado++ {

			fmt.Println(numBase, "x", multiplicado, "=", multiplicado*numBase)

		}
	}

}

func main() {

	salarioComBonus()
	carroOuMoto()
	taxaCarroOuMoto()
	descontoSalarial()
	descontoSalarial2()
	testantoFor()
	testandoBreak()
	testandoContinue()
	fazendoWhileComFor()
	forAninhados()
}

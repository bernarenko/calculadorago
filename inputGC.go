package main

import "fmt"

func inputQuerSubstituir(p1 int) (soun int) {
	fmt.Println("Você deseja fazer o grau C?")
	fmt.Println("1 - Sim")
	fmt.Println("2 - Não")
	fmt.Scanln(&p1)

	return p1
}

func inputSubsAouB(p1 int) (aoub int) {
	fmt.Println("Qual nota vocẽ substituirá?")
	fmt.Println("1 - Grau A")
	fmt.Println("2 - Grau B")
	fmt.Scanln(&p1)

	return p1
}

func inputProvaGC(p1 float32) (pgc float32) {
	fmt.Println("Insira sua nota de prova do grau C")
	fmt.Scanln(&p1)

	return p1
}

func inputSubstituindoTrabalhos(p1, p2 int) (nts int) {
	if p1 == 1 {
		fmt.Println("Você está substituindo o seu grau A")
		fmt.Println("Sua nota de trabalhos será substituida?")
		fmt.Scanln(&p2)
	}

	if p1 == 2 {
		fmt.Println("Você está substituindo o seu grau B")
		fmt.Println("Sua nota de trabalhos será substituida?")
		fmt.Println("1 - Sim")
		fmt.Println("2 - Não")
		fmt.Scanln(&p2)
	}

	return p2
}

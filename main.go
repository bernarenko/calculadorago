package main

import (
	"fmt"
)

func main() {
	var pga float32 // grau A
	var tga float32
	var pgb float32 // grau B
	var tgb float32
	mensagemInicial()
	pga, tga = calculoGA(pga, tga)
	pgb, tgb = calculoGB(pgb, tgb)

	var cga float32 = (((tga * 3) + (pga * 7)) / 10) // resultado do grau A.
	var cgb float32 = (((tgb * 3) + (pgb * 7)) / 10) // resultado do grau B.
	var pregc float32 = ((cga * 3.33) + (cgb * 6.67)) / 10

	mensagemNotaDeGrau(cga, cgb)

	if pregc >= 7 {
		fmt.Println("Você passou. Não recomendamos que faça o GC")
	} else {
		fmt.Println("Você não passou. Recomendamos que faça o GC")
	}

	// começa grau C aqui

	var soun int
	fmt.Println("Você deseja fazer o grau C?")
	fmt.Println("1 - Sim")
	fmt.Println("2 - Não")
	fmt.Scanln(&soun)
	if soun == 2 {
		if pregc >= 7 {
			fmt.Println("Ok! Parabéns! Aproveite as feŕias!")
		} else {
			fmt.Println("Tente estudar mais de uma próxima vez.")
		}
	}

	if soun == 1 {
		var aoub int
		fmt.Println("Qual nota vocẽ substituirá?")
		fmt.Println("1 - Grau A")
		fmt.Println("2 - Grau B")
		fmt.Scanln(&aoub)
		if aoub == 1 {
			var nts int
			fmt.Println("Você está substituindo o seu grau A")
			fmt.Println("Sua nota de trabalhos será substituida?")
			fmt.Scanln(&nts)
		}
		if aoub == 2 {
			var nts int
			fmt.Println("Você está substituindo o seu grau B")
			fmt.Println("Sua nota de trabalhos será substituida?")
			fmt.Println("1 - Sim")
			fmt.Println("2 - Não")
			fmt.Scanln(&nts)
			if nts == 2 {
				var pgc float32

				fmt.Println("Insira sua nota de prova do grau B")
				fmt.Scanln(&pgc)
				var posgc float32 = ((cga * 3.33) + (pgc * 6.67)) / 10

				if posgc >= 7 {
					fmt.Println("Você passou. Parabéns!")
				} else {
					fmt.Println("Você reprovou. Tente outra vez.")
				}
			}
		}
	}
}

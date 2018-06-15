package main

import (
	"fmt"
)

func main() {
	var pga, tga float32 // grau A
	var pgb, tgb float32 // grau B
	var cga, cgb, pregc float32
	var pgc float32
	var soun, aoub, nts int // Substituir GC || Qual substituir || Substituir notas
	var posgc float32
	var graucnotas float32

	mensagemInicial()
	// Inputs de dados de provas(pgx) e trabalhos(tgx).
	pga, tga = inputGA(pga, tga)
	pgb, tgb = inputGB(pgb, tgb)

	// Processamento de dados || cga = Grau A / cgb = Grau B || pregc = Calculo da Nota final, pré GC.
	cga = calculoDeGrau(tga, pga)
	cgb = calculoDeGrau(tgb, pgb)
	pregc = calculoGC(cga, cgb)

	mensagemNotaDeGrau(cga, cgb)
	mensagemResultadoPreGC(pregc)
	/////////////////////////////
	// começa grau C aqui //
	/////////////////////////////
	soun = inputQuerSubstituir(soun)
	if soun == 2 { // ao escolher não substituir
		mensagemNaoSubstituirGC(pregc)
	}
	if soun == 1 { // ao escolher substituir
		aoub = inputSubsAouB(aoub)                  // escolhendo qual grau será substituido
		nts = inputSubstituindoTrabalhos(aoub, nts) // escolhendo se trabalhos serão substituidos

		if aoub == 1 { // Substituindo grau A //
			if nts == 1 {
				pgc = inputProvaGC(pgc)
				posgc = calculoGC(pgc, cgb)

				if posgc >= 7 {
					fmt.Println("Você passou. Parabéns!")
				} else {
					fmt.Println("Você reprovou. Tente outra vez.")
				}
			}
			if nts == 2 {
				pgc = inputProvaGC(pgc)
				graucnotas = calculoDeGrau(tga, pgc)
				posgc = calculoGC(graucnotas, cgb)
				if posgc >= 7 {
					fmt.Println("Você passou. Parabéns!")
				} else {
					fmt.Println("Você reprovou. Tente outra vez.")
				}
			}
		}

		if aoub == 2 { // Substituindo grau B //
			if nts == 1 {
				pgc = inputProvaGC(pgc)
				posgc = calculoGC(cga, pgc)

				if posgc >= 7 {
					fmt.Println("Você passou. Parabéns!")
				} else {
					fmt.Println("Você reprovou. Tente outra vez.")
				}
			}

			if nts == 2 {
				pgc = inputProvaGC(pgc)
				graucnotas = calculoDeGrau(tgb, pgc)
				posgc = calculoGC(cga, graucnotas)
				if posgc >= 7 {
					fmt.Println("Você passou. Parabéns!")
				} else {
					fmt.Println("Você reprovou. Tente outra vez.")
				}
			}
		}
	}
}

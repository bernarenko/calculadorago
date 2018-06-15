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
	// Inputs de dados de provas(pgx) e trabalhos(tgx).
	pga, tga = inputGA(pga, tga)
	pgb, tgb = inputGB(pgb, tgb)

	// Processamento de dados || cga = Grau A / cgb = Grau B || pregc = Calculo da Nota final, pré GC.
	var cga float32 = calculoDeGrau(tga, pga)
	var cgb float32 = calculoDeGrau(tgb, pgb)
	var pregc float32 = calculoGC(cga, cgb)
	var pgc float32

	mensagemNotaDeGrau(cga, cgb)
	mensagemResultadoPreGC(pregc)

	// começa grau C aqui

	var soun int // inteiro para salvar opção de substituir ou não o grau C
	var aoub int // inteiro para salvar opção de qual será substituido
	var nts int  // inteiro para salvar opção de substituição || 1 - A, 2 -B

	/*	Onde começa a corrigir	*/
	soun = inputQuerSubstituir(soun)
	if soun == 2 { // ao escolher não substituir
		mensagemNaoSubstituirGC(pregc)
	}
	if soun == 1 { // ao escolher substituir
		aoub = inputSubsAouB(aoub)              // escolhendo Qual será substituido.
		nts = inputSubstituindoNotas(aoub, nts) // escolhendo se nota será substituida

		if aoub == 1 { // Substituindo grau A //
			pgc = inputProvaGC(pgc)
			var posgc float32 = calculoGC(pgc, cgb)

			if posgc >= 7 {
				fmt.Println("Você passou. Parabéns!")
			} else {
				fmt.Println("Você reprovou. Tente outra vez.")
			}
		}
		if aoub == 2 { // Substituindo grau B //
			if nts == 2 {
				pgc = inputProvaGC(pgc)
				var posgc float32 = calculoGC(cga, pgc)

				if posgc >= 7 {
					fmt.Println("Você passou. Parabéns!")
				} else {
					fmt.Println("Você reprovou. Tente outra vez.")
				}
			}
		}
	}
}

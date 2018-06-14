package main

import "fmt"

func mensagemInicial() {
	fmt.Println("Essa calculadora irá te ajudar com suas notas")
	fmt.Println("Vamos calcular seu GA, GB e GC.")

}

func mensagemNotaDeGrau(cga, cgb float32) {
	fmt.Println("\nResultado do Grau A:", cga)
	fmt.Println("Resultado do Grau B:", cgb)
	fmt.Printf("Essas foram suas notas\n\n\n")

}

func mensagemResultadoPreGC(pregc float32) {
	if pregc >= 7 {
		fmt.Println("Você passou. Não recomendamos que faça o GC")
	} else {
		fmt.Println("Você não passou. Recomendamos que faça o GC")
	}
}

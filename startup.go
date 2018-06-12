package main

import "fmt"

func mensagemInicial() {
	fmt.Println("Essa calculadora irá te ajudar com suas notas")
	fmt.Println("Vamos calcular seu GA, GB e GC.")

}

func calculoGA(p1, p2 float32) (pga, tga float32) {
	fmt.Println("\nInsira nota da sua prova do grau A:")
	fmt.Scanln(&p1)
	fmt.Println("Insira nota do seu trabalho do grau A:")
	fmt.Scanln(&p2)

	return p1, p2
}

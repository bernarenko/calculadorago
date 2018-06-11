package main

import "fmt"

func main() {
	var tga float32
	var pga float32

	fmt.Println("Vamos calcular suas notas Unisinos")
	fmt.Println("Insira nota da sua prova do grau A:")
	fmt.Scanln(&tga)
	fmt.Println("Insira nota do seu trabalho do grau A:")
	fmt.Scanln(&pga)
	var cga float32 = (((tga * 3) + (pga * 7)) / 10) // resultado do grau A.

	var tgb float32
	var pgb float32

	fmt.Println("Insira nota da sua prova do grau B:")
	fmt.Scanln(&tgb)
	fmt.Println("Insira nota do seu trabalho do grau B:")
	fmt.Scanln(&pgb)
	var cgb float32 = (((tgb * 3) + (pgb * 7)) / 10)

	var pregc float32 = ((cga * 3.33) + (cgb * 6.67)) / 10

	if pregc >= 7 {
		fmt.Println("Você passou. Não recomendamos que faça o GC")
	} else {
		fmt.Println("Você não passou. Recomendamos que faça o GC")
	}
}

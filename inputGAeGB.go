package main

import "fmt"

func inputGA(p1, p2 float32) (pga, tga float32) {
	fmt.Println("\nInsira nota da sua prova do grau A:")
	fmt.Scanln(&p1)
	fmt.Println("Insira nota do seu trabalho do grau A:")
	fmt.Scanln(&p2)

	return p1, p2
  }
func inputGB(p1, p2 float32) (pgb, tgb float32) {
	fmt.Println("Insira nota da sua prova do grau B:")
	fmt.Scanln(&p1)
	fmt.Println("Insira nota do seu trabalho do grau B:")
	fmt.Scanln(&p2)
	return p1, p2
}

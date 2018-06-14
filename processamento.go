package main

func calculoDeGrau(tgx, pgx float32) (cdg float32) {
	var resultadoDeGrau float32 = (((tgx * 3) + (pgx * 7)) / 10)

	return resultadoDeGrau
}

func calculoGC(cga, cgb float32) (pregc float32) {
	var resultado float32 = ((cga * 3.33) + (cgb * 6.67)) / 10

	return resultado
}

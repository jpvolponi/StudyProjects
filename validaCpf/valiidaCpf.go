package main

import "fmt"

func ValidaCpf() bool {
	var cpf string
	var sum, sum2, i, j, k, l, count int32
	var resp bool
	fmt.Println("\nDigite o CPF.")
	fmt.Scan(&cpf)
	//verifica primeiro dígito

	cpfInt, erro := ConvInt(cpf)
	if erro != nil {
		fmt.Println(erro)
	} else {
		for a := 0; a < len(cpfInt)-1; a++ {
			if cpfInt[a] == cpfInt[a+1] {
				count++
			}
		}
		if count >= 6 {
			resp = false
		} else {
			j = 10
			k = 11
			for i = 0; i < 9; i++ {
				sum += cpfInt[i] * j
				//fmt.Printf("DEBUG: %d * %d = %d\n", cpfInt[i], j, sum)
				j--
			}
			//verifica segundo digito
			for l = 0; l < 10; l++ {
				sum2 += cpfInt[l] * k
				//fmt.Printf("DEBUG: %d * %d = %d\n", cpfInt[l], k, sum2)
				k--
			}
		}
		primeiroDigito := sum * 10 % 11
		segundoDigito := sum2 * 10 % 11
		//fmt.Println(sum)
		//fmt.Println(primeiroDigito)
		//fmt.Println(segundoDigito)

		if primeiroDigito == cpfInt[9] && segundoDigito == cpfInt[10] {
			resp = true
		} else {
			resp = false
		}
	}
	return resp
}

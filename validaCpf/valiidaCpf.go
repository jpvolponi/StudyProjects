package main

import "fmt"

func CalculaPrimeiroDigito(cpf []int32) int32 {
	var sum int32
	var j int32 = 10
	for _, v := range cpf[0:9] {
		sum += v * j
		j--
	}
	primeiroDigito := sum * 10 % 11
	return primeiroDigito
}
func CalculaSegundoDigito(cpf []int32) int32 {
	var sum int32
	var j int32 = 11
	for _, v := range cpf[0:10] {
		sum += v * j
		j--
	}
	segundoDigito := sum * 10 % 11
	return segundoDigito
}
func FalsoValido(cpf []int32) bool {
	var count int32
	var resp bool = false
	for i := 0; i < len(cpf)-1; i++ {
		if cpf[i] == cpf[i+1] {
			count++
		}
	}
	if count >= 6 {
		resp = true
	}
	return resp
}

func ValidaCpf() bool {
	var cpf string
	var resp bool
	fmt.Println("\nDigite o CPF.")
	fmt.Scan(&cpf)

	//Converte string para silce de int32
	cpfInt, erro := ConvInt(cpf)
	if erro != nil {
		fmt.Println("Erro de conversão:", erro)
	} else if FalsoValido(cpfInt) {
		return false
	} else {
		//verifica digitos
		if CalculaPrimeiroDigito(cpfInt) == cpfInt[9] && CalculaSegundoDigito(cpfInt) == cpfInt[10] {
			resp = true
		} else {
			resp = false
		}
	}
	return resp
}

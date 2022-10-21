package main

import (
	"errors"
	"strconv"
	"strings"
)

func ConvInt(s string) ([]int32, error) {

	length := len(s)
	var int_split = make([]int32, length)
	st_split_teste := strings.Split(s, "")
	var erro error

	for i, v := range st_split_teste {
		if f, err := strconv.Atoi(v); err == nil {
			int_split[i] = int32(f)
		} else {
			erro = errors.New("Erro de Conversão.")
		}
	}
	if len(int_split) > 11 || len(int_split) < 11 {
		erro = errors.New("Quantidade de digitos insuficiente.")
	}
	return int_split, erro
}

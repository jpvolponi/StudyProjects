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
			erro = errors.New("Conversão errada")
		}
	}
	return int_split, erro
}

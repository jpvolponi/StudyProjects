package main

import "testing"

func TestFalsoValido(t *testing.T) {
	cpf := []int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	teste := FalsoValido(cpf)
	resultado := true
	if teste != resultado {
		t.Errorf("got %v, want %v", teste, resultado)
	}
}

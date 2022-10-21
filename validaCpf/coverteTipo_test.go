package main

import "testing"

func TestConvInt(t *testing.T) {
	_, erroTeste := ConvInt("12345678910")
	if erroTeste != nil {
		t.Errorf("got %v, want %v", erroTeste, nil)
	}
}

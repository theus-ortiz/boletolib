package boletolib

import "github.com/theus-ortiz/boletolib/banks/grafeno"

// Grafeno retorna a implementação do Banco 310 — Vórtx/Grafeno.
func Grafeno() *grafeno.Bank {
	return grafeno.New()
}

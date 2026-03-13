// Package grafeno implementa bank.Bank para o Banco 310 — Vórtx/Grafeno.
//
// Referência: "Grafeno_Vórtx CNAB 400 - Remessa" (Banco 310 - V1.0)
package grafeno

import (
	"fmt"
	"strings"

	"github.com/theus-ortiz/boletolib/internal/calc"
)

const (
	bankCode     = "310"
	currencyCode = "9"

	// wallet é o código da carteira usado como prefixo no cálculo do DV do Nosso Número.
	// Conforme documentação: "acrescentar o número da carteira (21) à esquerda
	// antes do Nosso Número, e aplicar o módulo 11 com base 7."
	wallet = "21"
)

// Bank implementa bank.Bank para o Banco 310 (Vórtx/Grafeno).
type Bank struct{}

// New retorna uma instância do banco Grafeno/Vórtx (Banco 310).
func New() *Bank {
	return &Bank{}
}

// Code retorna o código do banco: "310".
func (b *Bank) Code() string { return bankCode }

// CurrencyCode retorna o código da moeda: "9" (Real brasileiro).
func (b *Bank) CurrencyCode() string { return currencyCode }

// NossoNumeroDV calcula o dígito verificador do Nosso Número.
//
// Conforme documentação Grafeno:
//   - Prefixar o Nosso Número com a carteira "21"
//   - Aplicar Módulo 11 com base 7
//   - Se resto == 0 ou resultado == 10 ou 11 → dígito = 0
//
// Exemplos verificados:
//
//	NossoNumeroDV("00000000001") == "9"
//	NossoNumeroDV("00000000002") == "7"
func (b *Bank) NossoNumeroDV(nossoNumero string) string {
	nn := padLeft(nossoNumero, 11)
	input := wallet + nn
	dv := calc.Mod11Base7(input)
	return fmt.Sprintf("%d", dv)
}

// FreeField monta o campo livre de 25 dígitos conforme especificação Grafeno:
//
//	Posições 01–04 (4 dígitos) : Agência do Beneficiário
//	Posições 05–14 (10 dígitos): Conta do Beneficiário sem dígito verificador
//	Posições 15–25 (11 dígitos): Nosso Número sem dígito verificador
func (b *Bank) FreeField(agency, account, nossoNumero string) string {
	agencyStr  := padLeft(agency, 4)
	
	accountStr := padLeft(account, 10)
	nnStr      := padLeft(nossoNumero, 11)
	return agencyStr + accountStr + nnStr
}

// padLeft preenche a string com zeros à esquerda até atingir o tamanho desejado.
// Se a string já for maior que length, retorna os últimos `length` caracteres.
func padLeft(s string, length int) string {
	s = strings.TrimLeft(s, " ")
	if len(s) >= length {
		return s[len(s)-length:]
	}
	return strings.Repeat("0", length-len(s)) + s
}

package boletolib

import (
	"fmt"
	"strings"

	"github.com/theus-ortiz/boletolib/internal/calc"
)

// Generate gera o código de barras, linha digitável e Nosso Número com DV
// a partir dos dados do boleto.
//
// Fluxo de geração:
//  1. Valida os dados de entrada
//  2. Calcula o DV do Nosso Número (delegado ao banco)
//  3. Calcula o fator de vencimento (ciclo automático)
//  4. Formata o valor em centavos (10 dígitos)
//  5. Monta o campo livre de 25 dígitos (delegado ao banco)
//  6. Monta os 43 dígitos base do código de barras (sem DV)
//  7. Calcula o DV do código de barras (Módulo 11, base 9)
//  8. Monta o código de barras final (44 dígitos)
//  9. Monta a linha digitável (5 campos + DACs por Módulo 10)
func Generate(b Boleto) (Result, error) {
	if err := validate(b); err != nil {
		return Result{}, err
	}

	// Passo 1 — DV do Nosso Número
	dvNN := b.Bank.NossoNumeroDV(b.NossoNumero)

	// Passo 2 — Fator de vencimento (4 dígitos, seleção automática de ciclo)
	factor := fmt.Sprintf("%04d", calc.DueDateFactor(b.DueDate))

	// Passo 3 — Valor (10 dígitos em centavos, sem ponto nem vírgula)
	amount := fmt.Sprintf("%010d", b.AmountCents)

	// Passo 4 — Campo livre de 25 dígitos (composição definida pelo banco)
	freeField := b.Bank.FreeField(b.Agency, b.Account, b.NossoNumero)

	// Passo 5 — Base do código de barras (43 dígitos, sem a posição 5 do DV)
	//   Posições 1–4  : banco(3) + moeda(1)
	//   Posições 6–9  : fator(4)
	//   Posições 10–19: valor(10)
	//   Posições 20–44: campo livre(25)
	// Total: 3+1+4+10+25 = 43 dígitos
	barcodeBase := b.Bank.Code() + b.Bank.CurrencyCode() + factor + amount + freeField

	// Passo 6 — DV do código de barras (Módulo 11, base 9 sobre os 43 dígitos)
	dvBarcode := fmt.Sprintf("%d", calc.Mod11Base9(barcodeBase))

	// Passo 7 — Código de barras final: inserir DV na posição 5 (índice 4)
	barcode := barcodeBase[:4] + dvBarcode + barcodeBase[4:]

	// Passo 8 — Linha digitável
	typeable := buildTypeableLine(b, freeField, dvBarcode, factor, amount)

	return Result{
		Barcode:       barcode,
		TypeableLine:  typeable,
		NossoNumeroDV: b.NossoNumero + "-" + dvNN,
	}, nil
}

// buildTypeableLine monta os 5 campos da linha digitável.
//
// O campo livre de 25 dígitos é fatiado em 3 partes:
//
//	fatia 1: posições 01–05 → base do campo 1 (junto com banco e moeda)
//	fatia 2: posições 06–15 → base do campo 2
//	fatia 3: posições 16–25 → base do campo 3
//
// Formato final:
//
//	BBBMC.CCCCD CCCCCC.CCCCD CCCCCC.CCCCD D FFFFVVVVVVVVVV
func buildTypeableLine(b Boleto, freeField, dvBarcode, factor, amount string) string {
	fl1 := freeField[0:5]
	fl2 := freeField[5:15]
	fl3 := freeField[15:25]

	c1base := b.Bank.Code() + b.Bank.CurrencyCode() + fl1 // 9 dígitos
	c2base := fl2                                          // 10 dígitos
	c3base := fl3                                          // 10 dígitos

	dac1 := fmt.Sprintf("%d", calc.Mod10(c1base))
	dac2 := fmt.Sprintf("%d", calc.Mod10(c2base))
	dac3 := fmt.Sprintf("%d", calc.Mod10(c3base))

	campo1 := c1base[:5] + "." + c1base[5:] + dac1
	campo2 := c2base[:5] + "." + c2base[5:] + dac2
	campo3 := c3base[:5] + "." + c3base[5:] + dac3
	campo4 := dvBarcode
	campo5 := factor + amount

	return strings.Join([]string{campo1, campo2, campo3, campo4, campo5}, " ")
}

// validate verifica os dados de entrada antes da geração.
func validate(b Boleto) error {
	if b.Bank == nil {
		return ErrNilBank
	}
	if len(b.NossoNumero) != 11 {
		return ErrInvalidNossoNumero
	}
	if len(b.Agency) == 0 || len(b.Agency) > 4 {
		return ErrInvalidAgency
	}
	if len(b.Account) == 0 || len(b.Account) > 10 {
		return ErrInvalidAccount
	}
	if b.AmountCents < 0 {
		return ErrNegativeAmount
	}
	return nil
}

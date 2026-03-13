package boletolib

import (
	"errors"
	"time"

	"github.com/theus-ortiz/boletolib/bank"
)

// Boleto contém todos os dados necessários para gerar o código de barras.
type Boleto struct {
	Bank        bank.Bank // implementação do banco, ex: grafeno.New()
	Agency      string    // agência do beneficiário, somente dígitos
	Account     string    // conta do beneficiário sem dígito verificador, somente dígitos
	NossoNumero string    // 11 dígitos sem DV ("00000000001") ou com DV ("00000000001-9")
	DueDate     time.Time // data de vencimento
	Amount      float64   // valor em reais, ex: 1047.00 (= R$ 1.047,00)
}

// Result contém os dados gerados prontos para uso.
type Result struct {
	Barcode      string // 44 dígitos numéricos
	TypeableLine string // linha digitável formatada com pontos e espaços
	NossoNumeroDV string // nosso número com dígito: "NNNNNNNNNNN-D"
}

var (
	ErrInvalidNossoNumero    = errors.New("nosso número must be 11 digits, 12 digits (with DV), or \"NNNNNNNNNNN-D\"")
	ErrWrongNossoNumeroDV    = errors.New("nosso número DV provided is incorrect")
	ErrInvalidAgency         = errors.New("agency must be up to 4 digits")
	ErrInvalidAccount        = errors.New("account must be up to 10 digits")
	ErrInvalidAmount         = errors.New("amount must be >= 0")
	ErrNilBank               = errors.New("bank implementation must not be nil")
)

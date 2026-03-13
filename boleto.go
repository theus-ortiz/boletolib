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
	NossoNumero string    // 11 dígitos, sem dígito verificador
	DueDate     time.Time // data de vencimento
	AmountCents int64     // valor em centavos (ex: R$ 10,50 = 1050)
}

// Result contém os dados gerados prontos para uso.
type Result struct {
	Barcode      string // 44 dígitos numéricos
	TypeableLine string // linha digitável formatada com pontos e espaços
	NossoNumeroDV string // nosso número com dígito: "NNNNNNNNNNN-D"
}

var (
	ErrInvalidNossoNumero = errors.New("nosso número must be exactly 11 digits")
	ErrInvalidAgency      = errors.New("agency must be up to 4 digits")
	ErrInvalidAccount     = errors.New("account must be up to 10 digits")
	ErrNegativeAmount     = errors.New("amount in cents must be >= 0")
	ErrNilBank            = errors.New("bank implementation must not be nil")
)

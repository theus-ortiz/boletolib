package main

import (
	"fmt"
	"log"
	"time"

	boletolib "github.com/theus-ortiz/boletolib"
	"github.com/theus-ortiz/boletolib/banks/grafeno"
)

func main() {
	result, err := boletolib.Generate(boletolib.Boleto{
		Bank:        grafeno.New(),
		Agency:      "0001",
		Account:     "00110173",
		NossoNumero: "10000000005-6",
		DueDate:     time.Date(2026, 6, 22, 0, 0, 0, 0, time.UTC),
		Amount:      633.00,
	})
	if err != nil {
		log.Fatalf("erro ao gerar boleto: %v", err)
	}

	fmt.Printf("Código de barras : %s\n", result.Barcode)
	fmt.Printf("Linha digitável  : %s\n", result.TypeableLine)
	fmt.Printf("Nosso Número     : %s\n", result.NossoNumeroDV)
}

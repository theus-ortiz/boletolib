# boletolib

Biblioteca Go para geração de código de barras e linha digitável de boletos bancários brasileiros.

Suporte atual: **Banco 310 — Vórtx/Grafeno**.

## Instalação

```bash
go get github.com/theus-ortiz/boletolib
```

## Uso

```go
import (
    boletolib "github.com/theus-ortiz/boletolib"
    "github.com/theus-ortiz/boletolib/banks/grafeno"
)

result, err := boletolib.Generate(boletolib.Boleto{
    Bank:        grafeno.New(),
    Agency:      "0001",             // agência do beneficiário (até 4 dígitos)
    Account:     "0012345678",       // conta sem dígito verificador (até 10 dígitos)
    NossoNumero: "00000000001",      // 11 dígitos, sem dígito verificador
    DueDate:     time.Date(2026, 4, 7, 0, 0, 0, 0, time.UTC),
    AmountCents: 150000,             // valor em centavos — R$ 1.500,00
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(result.Barcode)      // 44 dígitos
fmt.Println(result.TypeableLine) // linha digitável formatada
fmt.Println(result.NossoNumeroDV) // ex: 00000000001-9
```

## Resultado

| Campo | Exemplo |
|---|---|
| `Barcode` | `31091140900001500000001000123456780000000001` |
| `TypeableLine` | `31090.00106 01234.567804 00000.000013 1 14090000150000` |
| `NossoNumeroDV` | `00000000001-9` |

## Fator de Vencimento

A lib seleciona automaticamente o ciclo correto:

| Ciclo | Base | Vigência |
|---|---|---|
| 1 | 07/10/1997 | até 21/02/2025 |
| 2 | 22/02/2025 | atual |

## Adicionando um novo banco

1. Crie o diretório `banks/<nome>/`
2. Implemente a interface `bank.Bank`:

```go
package <nome>

import "github.com/theus-ortiz/boletolib/internal/calc"

type Bank struct{}

func New() *Bank { return &Bank{} }

func (b *Bank) Code() string                                          { return "<codigo>" }
func (b *Bank) CurrencyCode() string                                  { return "9" }
func (b *Bank) NossoNumeroDV(nossoNumero string) string               { /* lógica do banco */ }
func (b *Bank) FreeField(agency, account, nossoNumero string) string  { /* layout do banco */ }
```

3. O `generator.go` **não muda**.

## Testes

```bash
go test ./...
```

## Licença

MIT

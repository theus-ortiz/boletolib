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
)

result, err := boletolib.Generate(boletolib.Boleto{
    Bank:        boletolib.Grafeno(), // Banco 310 — Vórtx/Grafeno
    Agency:      "0001",              // agência do beneficiário (até 4 dígitos)
    Account:     "0012345678",        // conta sem dígito verificador (até 10 dígitos)
    NossoNumero: "00000000001",       // veja formatos aceitos abaixo
    DueDate:     time.Date(2026, 4, 7, 0, 0, 0, 0, time.UTC),
    AmountCents: 150000,              // valor em centavos — R$ 1.500,00
})
if err != nil {
    log.Fatal(err)
}

fmt.Println(result.Barcode)      // 44 dígitos
fmt.Println(result.TypeableLine) // linha digitável formatada
fmt.Println(result.NossoNumeroDV) // ex: 00000000001-9
```

## Formatos do Nosso Número

| Formato | Exemplo | Comportamento |
|---|---|---|
| 11 dígitos | `"00000000001"` | DV calculado automaticamente |
| 12 dígitos | `"000000000019"` | DV fornecido (último dígito), validado |
| Com traço | `"00000000001-9"` | DV fornecido, validado |

Se o DV fornecido estiver errado, `Generate` retorna erro:
```
nosso número DV provided is incorrect: got "5", expected "9"
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

## Testes

```bash
go test ./...
```

## Licença

MIT

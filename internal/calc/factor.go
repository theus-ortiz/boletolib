package calc

import "time"

// O sistema bancário brasileiro usa dois ciclos para o fator de vencimento.
// Quando o fator atingiu 9999 no ciclo 1, iniciou-se um novo ciclo em 22/02/2025.
//
// Ciclo 1: base 07/10/1997 = fator 1000 | válido até 21/02/2025 (fator 9999)
// Ciclo 2: base 22/02/2025 = fator 1000 | vigente a partir desta data
var (
	cycle1Base = time.Date(1997, 10, 7, 0, 0, 0, 0, time.UTC)
	cycle1End  = time.Date(2025, 2, 21, 0, 0, 0, 0, time.UTC)
	cycle2Base = time.Date(2025, 2, 22, 0, 0, 0, 0, time.UTC)
)

const baseOffset = 1000

// DueDateFactor calcula o fator de vencimento de 4 dígitos.
// Seleciona automaticamente o ciclo correto baseado na data de vencimento.
//
// Fórmula:
//
//	fator = dias entre data_base e vencimento + 1000
func DueDateFactor(due time.Time) int {
	due = time.Date(due.Year(), due.Month(), due.Day(), 0, 0, 0, 0, time.UTC)

	var base time.Time
	if !due.After(cycle1End) {
		base = cycle1Base
	} else {
		base = cycle2Base
	}

	days := int(due.Sub(base).Hours() / 24)
	return days + baseOffset
}

package calc

// Mod10 calcula o dígito de auto-conferência (DAC) pelo Módulo 10.
//
// Algoritmo:
//  1. Percorre os dígitos da direita para a esquerda.
//  2. Multiplica alternando entre ×2 e ×1, começando com ×2 no dígito mais à direita.
//  3. Se o produto for ≥ 10, soma os dois dígitos do resultado (ex: 16 → 1+6 = 7).
//  4. Soma todos os resultados.
//  5. DAC = próximo múltiplo de 10 − soma. Se a soma já for múltiplo de 10, DAC = 0.
func Mod10(digits string) int {
	sum := 0
	double := true // começa multiplicando por 2 (da direita para a esquerda)

	for i := len(digits) - 1; i >= 0; i-- {
		n := int(digits[i] - '0')
		if double {
			n *= 2
			if n >= 10 {
				n = n/10 + n%10
			}
		}
		sum += n
		double = !double
	}

	remainder := sum % 10
	if remainder == 0 {
		return 0
	}
	return 10 - remainder
}

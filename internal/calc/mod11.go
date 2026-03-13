package calc

// Mod11Base7 calcula o dígito verificador pelo Módulo 11 com pesos 2–7.
//
// Usado para: dígito verificador do Nosso Número.
//
// Algoritmo:
//  1. Percorre os dígitos da direita para a esquerda.
//  2. Multiplica pelos pesos 2, 3, 4, 5, 6, 7, ciclando de volta a 2.
//  3. Soma os produtos.
//  4. resto = soma % 11
//  5. resultado = 11 − resto
//  6. Se resto == 0 OU resultado == 10 OU resultado == 11 → retorna 0.
func Mod11Base7(digits string) int {
	sum := 0
	weight := 2

	for i := len(digits) - 1; i >= 0; i-- {
		sum += int(digits[i]-'0') * weight
		weight++
		if weight > 7 {
			weight = 2
		}
	}

	remainder := sum % 11
	if remainder == 0 {
		return 0
	}
	result := 11 - remainder
	if result == 10 || result == 11 {
		return 0
	}
	return result
}

// Mod11Base9 calcula o dígito verificador pelo Módulo 11 com pesos 2–9.
//
// Usado para: dígito verificador do código de barras (posição 5 dos 44 dígitos).
//
// Algoritmo:
//  1. Percorre os dígitos da direita para a esquerda.
//  2. Multiplica pelos pesos 2, 3, 4, 5, 6, 7, 8, 9, ciclando de volta a 2.
//  3. Soma os produtos.
//  4. resto = soma % 11
//  5. resultado = 11 − resto
//  6. Se resultado == 0 OU resultado == 1 OU resultado > 9 → retorna 1.
func Mod11Base9(digits string) int {
	sum := 0
	weight := 2

	for i := len(digits) - 1; i >= 0; i-- {
		sum += int(digits[i]-'0') * weight
		weight++
		if weight > 9 {
			weight = 2
		}
	}

	remainder := sum % 11
	result := 11 - remainder
	if result == 0 || result == 1 || result > 9 {
		return 1
	}
	return result
}

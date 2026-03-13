package bank

// Bank define o contrato que cada instituição bancária deve implementar.
//
// Toda lógica específica do banco — campo livre, dígito verificador do
// Nosso Número, código do banco — fica encapsulada aqui.
// O gerador principal nunca conhece os detalhes de cada banco.
type Bank interface {
	// Code retorna o código do banco com 3 dígitos (ex: "310").
	Code() string

	// CurrencyCode retorna o código da moeda ("9" = Real).
	CurrencyCode() string

	// NossoNumeroDV calcula e retorna o dígito verificador do Nosso Número.
	// Recebe o Nosso Número puro com 11 dígitos e retorna apenas o dígito ("0"–"9").
	NossoNumeroDV(nossoNumero string) string

	// FreeField monta o campo livre de exatamente 25 dígitos.
	// Cada banco define sua própria composição conforme documentação do Bacen/banco.
	// agency e account já devem ser apenas dígitos; o banco é responsável
	// pelo zero-padding interno.
	FreeField(agency, account, nossoNumero string) string
}

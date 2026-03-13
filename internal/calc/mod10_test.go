package calc_test

import (
	"testing"

	"github.com/theus-ortiz/boletolib/internal/calc"
)

func TestMod10(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "campo 1 do exemplo do documento (banco 237, moeda 9, campo livre 1-5 = 00310)",
			input: "237900310",
			want:  2,
		},
		{
			name:  "campo 2 do exemplo do documento",
			input: "4003177200",
			want:  3,
		},
		{
			name:  "campo 3 do exemplo do documento",
			input: "2800952790",
			want:  5,
		},
		{
			name:  "resultado zero quando já é múltiplo de 10",
			input: "0000000000",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Mod10(tt.input)
			if got != tt.want {
				t.Errorf("Mod10(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

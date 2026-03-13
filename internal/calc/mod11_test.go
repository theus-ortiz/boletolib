package calc_test

import (
	"testing"

	"github.com/theus-ortiz/boletolib/internal/calc"
)

func TestMod11Base7(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Nosso Número 00000000001 → DV 9",
			input: "2100000000001",
			want:  9,
		},
		{
			name:  "Nosso Número 00000000002 → DV 7",
			input: "2100000000002",
			want:  7,
		},
		{
			name:  "resto 0 → DV 0",
			input: "0",
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Mod11Base7(tt.input)
			if got != tt.want {
				t.Errorf("Mod11Base7(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

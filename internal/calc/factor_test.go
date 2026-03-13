package calc_test

import (
	"testing"
	"time"

	"github.com/theus-ortiz/boletolib/internal/calc"
)

func TestDueDateFactor(t *testing.T) {
	tests := []struct {
		name string
		date string
		want int
	}{
		{
			name: "data base do ciclo 1 retorna fator 1000",
			date: "1997-10-07",
			want: 1000,
		},
		{
			// De 1997-10-07 a 2025-02-21 = 9999 dias → fator = 9999 + 1000 = 10999.
			// O comentário do guia ("8999 dias") está incorreto.
			name: "último dia antes do ciclo 2 (1997-10-07 + 9999 dias)",
			date: "2025-02-21",
			want: 10999,
		},
		{
			name: "primeiro dia do ciclo 2 retorna fator 1000",
			date: "2025-02-22",
			want: 1000,
		},
		{
			name: "16 dias após início do ciclo 2",
			date: "2025-03-10",
			want: 1016,
		},
		{
			name: "vencimento no mesmo dia que a data base do ciclo 1",
			date: "1997-10-08",
			want: 1001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			due, err := time.Parse("2006-01-02", tt.date)
			if err != nil {
				t.Fatalf("failed to parse date %q: %v", tt.date, err)
			}
			got := calc.DueDateFactor(due)
			if got != tt.want {
				t.Errorf("DueDateFactor(%s) = %d, want %d", tt.date, got, tt.want)
			}
		})
	}
}

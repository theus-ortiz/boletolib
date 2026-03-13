package grafeno_test

import (
	"testing"

	"github.com/theus-ortiz/boletolib/banks/grafeno"
)

func TestCode(t *testing.T) {
	b := grafeno.New()
	if got := b.Code(); got != "310" {
		t.Errorf("Code() = %q, want %q", got, "310")
	}
}

func TestCurrencyCode(t *testing.T) {
	b := grafeno.New()
	if got := b.CurrencyCode(); got != "9" {
		t.Errorf("CurrencyCode() = %q, want %q", got, "9")
	}
}

func TestNossoNumeroDV(t *testing.T) {
	b := grafeno.New()

	tests := []struct {
		nn   string
		want string
	}{
		{"00000000001", "9"}, // soma=13, 13%11=2, 11-2=9
		{"00000000002", "7"}, // soma=15, 15%11=4, 11-4=7
	}

	for _, tt := range tests {
		got := b.NossoNumeroDV(tt.nn)
		if got != tt.want {
			t.Errorf("NossoNumeroDV(%q) = %q, want %q", tt.nn, got, tt.want)
		}
	}
}

func TestFreeField(t *testing.T) {
	b := grafeno.New()

	agency      := "0031"
	account     := "7720000"
	nossoNumero := "28009527900"

	fl := b.FreeField(agency, account, nossoNumero)

	if len(fl) != 25 {
		t.Fatalf("FreeField length = %d, want 25", len(fl))
	}

	gotAgency  := fl[0:4]
	gotAccount := fl[4:14]
	gotNN      := fl[14:25]

	wantAgency  := "0031"
	wantAccount := "0007720000"
	wantNN      := "28009527900"

	if gotAgency != wantAgency {
		t.Errorf("agency in free field = %q, want %q", gotAgency, wantAgency)
	}
	if gotAccount != wantAccount {
		t.Errorf("account in free field = %q, want %q", gotAccount, wantAccount)
	}
	if gotNN != wantNN {
		t.Errorf("nosso número in free field = %q, want %q", gotNN, wantNN)
	}
}

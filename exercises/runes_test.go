package main

import "testing"

const (
	A = "0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;"
)

func Example() {
	main()
	// Output:
	// Please provide one or more words to search
}

func TestParsingLine(t *testing.T) {
	char, name := ParseLine(A)
	charDesejado := 'A'
	if char != charDesejado {
		t.Errorf("Char desejado: %q | Recebido: %q", charDesejado, char)
	}

	nameDesejado := "LATIN CAPITAL LETTER A"
	if name != nameDesejado {
		t.Errorf("Name desejado: %q | Recebido: %q", nameDesejado, name)
	}
}

// t.Errorf: reporta o erro
// t.Fail : acusa o erro e encerra a execução

package ast

import (
	"testing"

	"github.com/cian911/platoon/token"
)

func TestString(t *testing.T) {
	input := "let myVar = anotherVar;"

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != input {
		t.Errorf("program.String() did not have the correct value. Want %v, Got %v", input, program.String())
	}
}

package lex

import (
	"fmt"
	"testing"
)

// Global lexer
var testLexer *Lexer = nil

func TestInit(t *testing.T) {
	var exprs, err = InitLexExprs([][]string{
		{"[0-9]+", "INT"},
		{"\"(.*?)\"", "STRING"},
		{"[ |\n|\t]+", "IGNORE"},
	})

	if err != nil {
		t.Errorf("Initialization error: %s", err)
	}

	t.Logf("[STATUS] :: SUCCESS\n")
	testLexer = InitLexer("IGNORE", exprs)
}

type TestCase struct {
	arg  string
	want []Token
}

func TestLexerPositive(t *testing.T) {
	cases := []TestCase{
		{
			arg: " 123 ",
			want: []Token{
				{Value: "123", Tag: "INT"},
			},
		},
		{
			arg: " \"meow\" 123 ",
			want: []Token{
				{Value: `"meow"`, Tag: "STRING"},
				{Value: "123", Tag: "INT"},
			},
		},
		{
			arg: " \n \t \"meow\" 12\n3 ",
			want: []Token{
				{Value: `"meow"`, Tag: "STRING"},
				{Value: "12", Tag: "INT"},
				{Value: "3", Tag: "INT"},
			},
		},
	}

	for c, tc := range cases {
		got, err := testLexer.Lex(tc.arg)
		want := tc.want
		t.Logf("Running test-case #%d :: arg:\"%s\", want:%s\n", c, tc.arg, want)

		if fmt.Sprint(got) != fmt.Sprint(want) {
			t.Errorf("got %s, want %s, err %s", got, want, err)
		} else {
			t.Logf("[STATUS] :: SUCCESS\n")
		}
	}
}

package lex

import (
	"errors"
	"fmt"
	"regexp"
)

type Token struct {
	Value string
	Tag   string
}

func (this *Token) String() string {
	return fmt.Sprintf("<Token: val: %s | tag: %s>", this.Value, this.Tag)
}

type Expression struct {
	matcher *regexp.Regexp
	tag     string
}

type Lexer struct {
	Expressions   []Expression
	IgnoreTokName string
}

func InitLexExprs(exprs [][]string) (tokens []Expression, err error) {
	for _, tok := range exprs {
		regex, tokName := tok[0], tok[1]
		matcher, err := regexp.Compile(regex)

		if err != nil {
			return nil, err
		}

		tokens = append(
			tokens, Expression{matcher: matcher, tag: tokName})
	}

	return
}

func InitLexer(ignore string, exprs []Expression) *Lexer {
	lex := new(Lexer)
	*lex = Lexer{Expressions: exprs, IgnoreTokName: ignore}
	return lex
}

func isTokAtStart(in string, matcher *regexp.Regexp) bool {
	return matcher.FindStringIndex(in) != nil && matcher.FindStringIndex(in)[0] == 0
}

func tryMatch(in string, matcher *regexp.Regexp) (string, int, bool) {
	if isTokAtStart(in, matcher) {
		value := matcher.FindString(in)
		nextPos := matcher.FindStringIndex(in)[1]

		return value, nextPos, true
	}
	return "", 0, false
}

func throwIllegalTokenErr(in string, pos int) error {
	var errorMsg string
	afterPosLen := len(in[pos:])

	if afterPosLen >= 3 {
		afterPosLen = 3
	}

	errorMsg = fmt.Sprintf(
		"Illegal token at position: %d, %s...", pos, in[pos:afterPosLen])
	return errors.New(errorMsg)
}

func (this *Lexer) Lex(in string) (tokens []Token, err error) {
	position := 0

	for position < len(in) {
		in = in[position:]
		match := false

		for _, exp := range this.Expressions {
			matcher, tag := exp.matcher, exp.tag
			value, nextPos, isMatched := tryMatch(in, matcher)

			if isMatched {
				match = true
				position = nextPos
				if tag != this.IgnoreTokName {
					tokens = append(tokens, Token{value, tag})
				}
				break
			}
			match = false
		}

		if !match {
			return nil, throwIllegalTokenErr(in, position)
		}
	}
	return tokens, nil
}

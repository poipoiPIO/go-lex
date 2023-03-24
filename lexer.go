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
	Matcher *regexp.Regexp
	Tag     string
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
			tokens, Expression{Matcher: matcher, Tag: tokName})
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

func (this *Lexer) Lex(in string) (tokens []Token, err error) {
	position := 0

	for position < len(in) {
		in = in[position:]
		match := false
		nextTokPos := 0

		for _, exp := range this.Expressions {
			matcher, tag := exp.Matcher, exp.Tag

			if isTokAtStart(in, matcher) {
				match = true
				nextTokPos = matcher.FindStringIndex(in)[1]

				if tag != this.IgnoreTokName {
					tokens = append(tokens, Token{
						Value: matcher.FindString(in),
						Tag:   tag,
					})
				}

				break
			}
			match = false
		}

		if !match {
			err = errors.New(fmt.Sprintf("Illegal token at position: %d, %s", position, in[position:position+3]))
			tokens = nil
			return
		}

		position = nextTokPos
	}

	return tokens, nil
}

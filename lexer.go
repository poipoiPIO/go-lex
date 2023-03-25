package lex

import (
	"errors"
	"fmt"
	"regexp"
)

// Token type represents the lexing result's token.
// Value : string with the resulting value matched by the lexer.
// Tag : the name of tag of the token represented as string.
type Token struct {
	Value string
	Tag   string
}

func (this *Token) String() string {
	return fmt.Sprintf("<Token: val: %s | tag: %s>", this.Value, this.Tag)
}

type expression struct {
	matcher *regexp.Regexp
	tag     string
}

type Lexer struct {
	Expressions   []expression
	IgnoreTokName string
}

func initLexExprs(exprs [][]string) (tokens []expression, err error) {
	for _, tok := range exprs {
		regex, tokName := tok[0], tok[1]
		matcher, err := regexp.Compile(regex)

		if err != nil {
			return nil, err
		}

		tokens = append(
			tokens, expression{matcher: matcher, tag: tokName})
	}

	return
}

// Lexer initialization function, provides human interface for
// creating lexers. it returns the new Lexer instance
// or error in case of an error.
// rules : slice of pairs { "regexp", "token tag name" } representing
// lexing rules of new lexer instance.
// ignore: the name of tag to be ignored by the lexer.
func InitLexer(rules [][]string, ignore string) (*Lexer, error) {
  e, err := initLexExprs(rules)

  if err != nil {
    return nil, err
  }

  lex := Lexer{Expressions: e, IgnoreTokName: ignore}
	return &lex, nil
}

func isTokAtStart(in string, matcher *regexp.Regexp) bool {
	return matcher.FindStringIndex(in) != nil &&
    matcher.FindStringIndex(in)[0] == 0
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

// Method that provides lexing facilities for the string.
// Returns the slice of tokens as the result of lexing
// or error in case of something went wrong.
// in: string representing the lexer's input
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

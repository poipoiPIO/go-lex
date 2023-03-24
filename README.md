<h2 align="center">
  Go-lex : Simple zero-dependency universal lexing library for Golang</br>
</h2>

![](https://github.com/poipoiPIO/go-lex/actions/workflows/on-push.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/poipoiPIO/go-lex.svg)](https://pkg.go.dev/github.com/poipoiPIO/go-lex)

## Example:
```go
func main() {
  // Lexing rules are represented as the slice of
  // Pairs that looks like: {"regex string", "token tag"}
  var exprs, _ = InitLexExprs([][]string{
    {"[0-9]+", "INT"},
    {"[ |\n|\t]+", "IGNORE"},
  })
  
  // Lexer initialization function takes two arguments
  // First argument:  ignoring token tag
  // Second argument: Lexing rules output of InitLexExprs()
  lexer := InitLexer("IGNORE", exprs)
  
  somethingToLex := "   123 \n123" // an string to lex
  
  // Lexer.Lex() function doing lexing
  // For provided string, and returns result
  // In form of Lex.Token slice, or error
  // In case of failure
  var out, _ = lexer.Lex(somethingToLex)
  
  // out is:
  //   []Token{{"123", "INT"}, {"123", "INT"}}
}
```

## API Reference:
### Data structures:
##### Lexer:
> type Lexer struct { Expressions []Expression, IgnoreTokName string }

Struct representing lexer type of library.

##### Expression:
> type Expression struct { matcher *Regexp, tag string }

Struct representing type of lexing rules for the `Lexer` purposes.

##### Token:
> type Token struct { Value string, Tag string }

Token struct are using to represent the result of lexing in user-friendly manier.

### Functions and methods:
#### InitLexExprs:
> func InitLexExprs(rules [][]string) ([]Expression, error)

Function needed to initialize lexing rules before lexer initialization.
It takes rules in form of the {"Rule Regexp", "Tag name"} pair slice and
returns the slice of rules for lexer.

#### InitLexer:
> func InitLexExprs(ignoreTag string, rules []Expression) *Lexer

Function needed to initialize lexer.
It takes the tag ignored to lexer and lexing rules from `InitLexExprs` function.
And returns the instance of `Lexer` struct needed to provide lexing service.

#### Lexer.Lex:
> func (*Lexer) Lex(input string) ([]Token, error)

The main lexing function of library. It provides lexing facilities for the
the provided string, and returns error or the slice of resulting `Tokens` in
case of success.

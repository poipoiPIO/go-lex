<h2 align="center">
  Go-lex : Simple zero-dependency universal lexing library for Golang</br>
</h2>

![](https://github.com/poipoiPIO/go-lex/actions/workflows/on-push.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/poipoiPIO/go-lex.svg)](https://pkg.go.dev/github.com/poipoiPIO/go-lex)

## Example:
```go
func main() {
  var lexer, _ = InitLexer([][]string{
    {"[0-9]+", "INT"},
    {"[ |\n|\t]+", "IGNORE"},
  }, "IGNORE")
  
  somethingToLex := "   123 \n123"
  
  var out, _ = lexer.Lex(somethingToLex)
  
  // out is:
  //   []Token{{"123", "INT"}, {"123", "INT"}}
}
```

## API Reference:
### Data structures:
##### Lexer:
> type Lexer struct { Expressions []expression, IgnoreTokName string }

Struct representing lexer type of library.

##### Token:
> type Token struct { Value string, Tag string }

Token struct are using to represent the result of lexing in user-friendly manier.

### Functions and methods:
#### InitLexer:
> func InitLexExprs(rules [][]string, ignoreTag string) (*Lexer, error)

Function needed to initialize lexer.
It takes the tag ignored to lexer and lexing rules 
represented as the slice of pairs `{ "regexp", "token tag" }`.
And returns the instance of `Lexer` struct needed to provide lexing service 
or the error in case of an error.

#### Lexer.Lex:
> func (*Lexer) Lex(input string) ([]Token, error)

The main lexing function of library. It provides lexing facilities for the
the provided string, and returns error or the slice of resulting `Tokens` in
case of success.

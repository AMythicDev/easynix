package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestParems(t *testing.T) {
	is := is.New(t)

	s := strings.NewReader("(")
	scanner := NewScanner(s)
	tok, err := scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, OpenParen)
	is.Equal(tok.line, int32(1))
	is.Equal(tok.col, int32(1))

	s = strings.NewReader(")")
	scanner = NewScanner(s)
	tok, err = scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, CloseParen)
	is.Equal(tok.line, int32(1))
	is.Equal(tok.col, int32(1))
}

func TestSymbols(t *testing.T) {
	is := is.New(t)

	s := strings.NewReader("+")
	scanner := NewScanner(s)
	tok, err := scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, Ident)
	is.Equal(tok.word, "+")

	s = strings.NewReader("-")
	scanner = NewScanner(s)
	tok, err = scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, Ident)
	is.Equal(tok.word, "-")

	s = strings.NewReader("*")
	scanner = NewScanner(s)
	tok, err = scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, Ident)
	is.Equal(tok.word, "*")

	s = strings.NewReader("/")
	scanner = NewScanner(s)
	tok, err = scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, Ident)
	is.Equal(tok.word, "/")
}

func TestBasicArithmetic(t *testing.T) {
	is := is.New(t)
	s := strings.NewReader("(+ 5 6)")
	scanner := NewScanner(s)

	tokens := [...]Token{
		{ttype: OpenParen, line: 1, col: 1},
		{ttype: Ident, line: 1, col: 2, word: "+"},
		{ttype: Numeric, line: 1, col: 4, word: "5"},
		{ttype: Numeric, line: 1, col: 6, word: "6"},
		{ttype: CloseParen, line: 1, col: 7},
		{ttype: EOF, line: 1, col: 8},
	}
	toks, err := scanner.readTillEOF()
	is.NoErr(err)
	is.True(slices.Equal(tokens[0:], toks))
}

func TestArithmeticNested(t *testing.T) {
	is := is.New(t)
	s := strings.NewReader("(* (- (+ 5 6) 10) 2)")
	scanner := NewScanner(s)

	tokens := [...]Token{
		{ttype: OpenParen, line: 1, col: 1},
		{ttype: Ident, line: 1, col: 2, word: "*"},
		{ttype: OpenParen, line: 1, col: 4},
		{ttype: Ident, line: 1, col: 5, word: "-"},
		{ttype: OpenParen, line: 1, col: 7},
		{ttype: Ident, line: 1, col: 8, word: "+"},
		{ttype: Numeric, line: 1, col: 10, word: "5"},
		{ttype: Numeric, line: 1, col: 12, word: "6"},
		{ttype: CloseParen, line: 1, col: 13},
		{ttype: Numeric, line: 1, col: 15, word: "10"},
		{ttype: CloseParen, line: 1, col: 17},
		{ttype: Numeric, line: 1, col: 19, word: "2"},
		{ttype: CloseParen, line: 1, col: 20},
		{ttype: EOF, line: 1, col: 21},
	}
	toks, err := scanner.readTillEOF()
	is.NoErr(err)
	is.True(slices.Equal(tokens[0:], toks))
}

func TestIdents(t *testing.T) {
	is := is.New(t)

	s := strings.NewReader("myvar")
	scanner := NewScanner(s)
	tok, err := scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, Ident)
	is.Equal(tok.word, "myvar")
}

func TestKeywords(t *testing.T) {
	is := is.New(t)

	s := strings.NewReader("defvar")
	scanner := NewScanner(s)
	tok, err := scanner.readToken()
	is.NoErr(err)
	is.Equal(tok.ttype, Keyword)
	is.Equal(tok.word, "defvar")
}

func TestStatment(t *testing.T) {
	is := is.New(t)

	s := strings.NewReader("(defvar x 5)")
	scanner := NewScanner(s)

	tokens := [...]Token{
		{ttype: OpenParen, line: 1, col: 1},
		{ttype: Keyword, line: 1, col: 2, word: "defvar"},
		{ttype: Ident, line: 1, col: 9, word: "x"},
		{ttype: Numeric, line: 1, col: 11, word: "5"},
		{ttype: CloseParen, line: 1, col: 12},
		{ttype: EOF, line: 1, col: 13},
	}
	toks, err := scanner.readTillEOF()
	is.NoErr(err)
	is.True(slices.Equal(tokens[0:], toks))
}

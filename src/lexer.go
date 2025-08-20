package main

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int8

const (
	OpenParen TokenType = iota
	Keyword
	Ident
	Numeric
	CloseParen
	EOF
)

var keywords = map[string]struct{}{
	"defvar": {},
}

type Scanner struct {
	src       bufio.Reader
	line, col int32
	pos       int64
}

type Token struct {
	ttype TokenType
	word  string
	line  int32
	col   int32
}

func NewScanner(src io.Reader) Scanner {
	return Scanner{
		src:  *bufio.NewReader(src),
		line: 1,
		col:  1,
		pos:  1,
	}
}

func (s *Scanner) readTillEOF() ([]Token, error) {
	var tok Token
	var err error
	token_list := make([]Token, 0)

	for tok.ttype != EOF {
		tok, err = s.readToken()
		if err != nil {
			return token_list, err
		}
		token_list = append(token_list, tok)
	}
	return token_list, nil
}

func (s *Scanner) readToken() (Token, error) {
	s.skipWhitespace()
	r, err := s.peekRune()
	if err == io.EOF {
		return Token{ttype: EOF, line: s.line, col: s.col}, nil
	}
	if err != nil {
		return Token{}, fmt.Errorf("cannot read byte")
	}

	switch r {
	case '(':
		tok := Token{
			ttype: OpenParen,
			line:  s.line,
			col:   s.col,
		}
		s.readRune()
		return tok, nil
	case ')':
		tok := Token{
			ttype: CloseParen,
			line:  s.line,
			col:   s.col,
		}
		s.readRune()
		return tok, nil
	default:
		tok := Token{
			ttype: -1,
			line:  s.line,
			col:   s.col,
		}

		if unicode.IsLetter(r) || isAllowedSymbol(r) {
			res, err := s.readIdent()
			if err != nil {
				return Token{}, err
			}
			if _, isKw := keywords[res]; isKw {
				tok.ttype = Keyword
			} else {
				tok.ttype = Ident
			}
			tok.word = res
		} else if unicode.IsDigit(r) {
			res, err := s.readNumeric()
			if err != nil {
				return Token{}, err
			}
			tok.ttype = Numeric
			tok.word = res
		}
		if tok.ttype == -1 {
			return Token{}, fmt.Errorf("invalid token found")
		}
		return tok, nil
	}
}

func (s *Scanner) readIdent() (string, error) {
	r, err := s.readRune()
	if err != nil {
		return "", nil
	}
	if !unicode.IsLetter(r) && !isAllowedSymbol(r) {
		return "", fmt.Errorf("invalid character at identifier position")
	}

	var result strings.Builder
	result.WriteRune(r)

	r, err = s.peekRune()

	for err == nil {
		if unicode.IsSpace(r) {
			return result.String(), nil
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !isAllowedSymbol(r) {
			return "", fmt.Errorf("invalid character at identifier position")
		}
		result.WriteRune(r)
		s.readRune()
		r, err = s.peekRune()
	}
	return result.String(), nil
}

func (s *Scanner) readNumeric() (string, error) {
	r, err := s.readRune()
	if err != nil {
		return "", nil
	}
	if !unicode.IsDigit(r) && r != '-' && r != '+' && r != '.' {
		return "", fmt.Errorf("invalid character at identifier position")
	}

	var result strings.Builder
	result.WriteRune(r)

	r, err = s.peekRune()

	for err == nil {
		if unicode.IsSpace(r) || r == ')' {
			return result.String(), nil
		}
		if !unicode.IsDigit(r) && r != '-' && r != '+' && r != '.' {
			return "", fmt.Errorf("invalid character at identifier position")
		}
		result.WriteRune(r)
		s.readRune()
		r, err = s.peekRune()
	}
	return result.String(), nil
}

func (s *Scanner) skipWhitespace() {
	c, err := s.peekRune()
	for err == nil {
		switch c {
		case ' ':
		case '\t':
		case '\n':
			s.line += 1
		default:
			return
		}
		s.readRune()
		c, err = s.peekRune()
	}
}

func (s *Scanner) peekRune() (rune, error) {
	b, err := s.src.Peek(4)
	if (len(b) == 0 && err == io.EOF) || (err != nil && err != io.EOF) {
		return -1, err
	}
	r, _ := utf8.DecodeRune(b)
	return r, nil
}

func (s *Scanner) readRune() (rune, error) {
	r, _, err := s.src.ReadRune()
	if err != nil {
		return 0x0, err
	}
	s.pos += 1
	s.col += 1
	return r, nil
}

func isAllowedSymbol(r rune) bool {
	var disabledSymbols = [...]rune{';'}
	return unicode.IsSymbol(r) || unicode.IsPunct(r) && r != ';' && !slices.Contains(disabledSymbols[0:], r)
}

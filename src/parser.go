package main

type Parser struct {
	scanner Scanner
	// currToken holds the current token being parsed. Prefer using the
	// [Parser.peekToken] or [Parser.consumeToken] method over accessing this
	// directly
	currToken Token
}

func NewParser(s Scanner) Parser {
	return Parser{
		scanner:   s,
		currToken: Token{ttype: -1},
	}
}

func (p *Parser) ParseProgram() (Token, error) {
	tok, err := p.peekToken()

	if err == nil {
		return Token{}, err
	}

	switch tok.ttype {
	case OpenParen:
		p.consumeToken()
		return p.parseSList()
	}
}

func (p *Parser) parseSList() (Token, error) {

}

func (p *Parser) peekToken() (Token, error) {
	if p.currToken.ttype == -1 {
		tok, err := p.scanner.readToken()
		if err != nil {
			return Token{ttype: -1}, nil
		}
		p.currToken = tok
	}
	return p.currToken, nil
}

func (p *Parser) consumeToken() (Token, error) {
	tok, err := p.scanner.readToken()
	if err != nil {
		return Token{ttype: -1}, nil
	}
	p.currToken = tok
	return p.currToken, nil
}

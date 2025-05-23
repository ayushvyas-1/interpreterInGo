package lexer

import "ayush.interpreter.monkey/src/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte // current character/byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // creates Lexer struct and initiates input from test input.
	l.readChar()              // reads and saves the current character to ch and moves position to readposition and readposition to +1
	return l                  // returns the struct reference.
}

/*
	this newToken() function takes token Types and current character

and returns a Token { Type: type from argument,

	  Literal: that character from argument
	}
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/*
this NextToken look at the current ch and return its associated token.
and also **we advance our pointer of input.
so when we call NextToken() again l.ch field is already updated.
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()         // returns the identifier...like the whole word:-> "fn" or "let"
			tok.Type = token.LooupIdent(tok.Literal) // matches the identifier and returns the Type. "FUNCTION" or "LET"
			return tok
		} else if isDigit(l.ch) { // handle if a digit is parsed. like 5 in 'let five = 5;'
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

/*
This peekChar() func returns the very next character
*/
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

/*
This func saves the numbers to lexer's ch field and retruns that as string..
*/
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

/*
this function reads Identifiers **also advances lexers positions until it encounters nonLetter
*/
func (l *Lexer) readIdentifier() string {
	position := l.position // takes position from

	for isLetter(l.ch) {
		l.readChar() // stores identifiers.
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) { // if readPosition is greater than length of input -> get it to match EOF
		l.ch = 0 // 0 so it matches with EOF.
	} else {
		l.ch = l.input[l.readPosition] /// else save the next letter(character) to lexer's ch property.
	}
	l.position = l.readPosition // move position to readPosition so
	l.readPosition += 1         // move readPosition ahead.
}

/*this function skips spaces..and new lines
 */
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

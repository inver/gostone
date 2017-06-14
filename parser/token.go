package parser

type TokenType int

type Token struct {
	Value     string
	Types     []TokenType
	Index     int
	IsIgnored bool
}

func NewTokenSingleType(value string, tType TokenType, index int) Token {
	return Token{value, append(make([]TokenType, 0), tType), index, false}
}

func NewTokenMultiTypes(value string, types []TokenType, index int) Token {
	return Token{value, append(make([]TokenType, 0), types...), index, false}
}

const (
	PROP          TokenType = iota + 1
	STATEMENT
	HEX
	FLOAT
	INT
	REF
	SPACES
	ID
	EOL
	NULL
	TRUE
	FALSE
	STAR
	PLUS
	DOT
	AND
	MINUS
	OR
	COLON
	COMMA
	GT
	LT
	LE
	GE
	LPAREN
	RPAREN
	LBRACKET
	RBRACKET
	NEQ
	EQ
	MOD
	NOT
	QUERY
	SLASH
	DQUOTE
	SQUOTE
	BACKSLASH
	METHOD
	BLOCK_START
	BLOCK_END
	CMT_START
	CMT_END
	LITERAL_START
	LITERAL_END
	THIS
	GLOBAL
	ELSEIF
	RETURN
	MACRO
	ELSE
	IF
	IN
	FOR
	VAR
	BOR
	BXOR
	BAND
	BIN
	SUPPRESS
	LISTEN
	TRIGGER
	ARROW
	AS
	BREAK
	CONTINUE
	WHILE
	AST_START
	AST_END

	EOF   TokenType = -1
	ERROR TokenType = -2
	TOKEN TokenType = -3
)

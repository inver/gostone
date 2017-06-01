package common

type AstType int

const (
	NOP     AstType = iota
	ARRAY
	REGEXP
	THIS
	GLOBAL
	NOT
	AND
	OR
	TERNARY
	ADD
	SUB
	MUL
	DIV
	MOD
	USUB
	LT
	GT
	LE
	GE
	EQ
	NEQ
	REF
	CALL
	VAR
	IF
	FOR
	MACRO
	// {{if ...}} {{return 10}} {{/if}} - throws from the template
	// {{for ...}} {{return 10}} {{/for}} - throws from the template
	// {{macro ...}} {{return 10}} {{/macro}} - throws from the macro
	// {{[foo: {{ bla bla {{return 10}} }} ] }} - throws from the block of code
	RETURN
	NODES
	NODELIST

	BOR
	BXOR
	BAND

	SUPPRESS
	BLS
	BRS
	BNOT

	BREAK
	CONTINUE
	WHILE

	T_NOP      AstType = -1
	T_BREAK    AstType = -2
	T_ARRAY    AstType = -3
	VALUE_NODE AstType = -100500
)

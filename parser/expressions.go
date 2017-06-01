package parser

type Expression struct {
	Expression string
	Types      []Token
}

var Expressions = [...]Expression{
	{"null\\b", []Token{PROP, NULL}},
	{"true\b", []Token{PROP, TRUE}},
}

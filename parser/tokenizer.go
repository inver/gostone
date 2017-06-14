package parser

import (
	"regexp"
	"strings"
)

type Tokenizer struct {
	inputLen  int
	input     string
	baseUri   string
	matcher   *regexp.Regexp
	lastIndex int
	buffer    []Token
}

type TokenizerResult struct {
	Tokens  []Token
	IsFound bool
}

func NotFoundTokenizerResult() TokenizerResult {
	return TokenizerResult{nil, false}
}

func FoundTokenizerResult() TokenizerResult {
	return TokenizerResult{nil, true}
}

func NewTokenizerResult(tokens ...Token) TokenizerResult {
	return TokenizerResult{tokens, true}
}

func (tr *TokenizerResult) Value() string {
	return tr.Tokens[0].Value
}

func (tr *TokenizerResult) First() Token {
	return tr.Tokens[0]
}

func NewTokenizer(input, baseUri string) *Tokenizer {
	t := new(Tokenizer)
	t.input = input
	t.inputLen = len(input)
	t.baseUri = baseUri
	t.lastIndex = 0
	t.matcher = getRegexp()
	return t
}

func getRegexp() *regexp.Regexp {
	tokens := make([]string, len(Expressions))
	for i, expr := range Expressions {
		tokens[i] = "(" + expr.Expression + ")"
		//tokens[i] = expr.Expression
	}
	str := strings.Join(tokens, "|")
	r, err := regexp.Compile(str)
	if err != nil {
		panic(err)
	}
	return r
}

func (t *Tokenizer) readTokenToBuffer() {
	checkIndex := t.inputLen;
	if (t.lastIndex >= checkIndex) {
		// return T_EOF if we reached end of file
		t.buffer = append(t.buffer, Token{"", append(make([]TokenType, 0), EOF), checkIndex, false})
		return
	}

	input := t.input[t.lastIndex:]
	indexes := t.matcher.FindStringSubmatchIndex(input)

	if len(indexes) > 0 {
		if indexes[0] > t.lastIndex {
			t.buffer = append(t.buffer, NewTokenSingleType(input[t.lastIndex:indexes[0]], ERROR, t.lastIndex))
		}

		submatches := t.matcher.FindStringSubmatch(input)

		i := 1
		for ; i < len(submatches) && submatches[i] == ""; {
			i++
		}

		t.buffer = append(t.buffer, NewTokenMultiTypes(submatches[i], Expressions[i-1].Types, t.lastIndex))
		t.lastIndex = t.lastIndex + indexes[1]
	} else {
		t.buffer = append(t.buffer, NewTokenSingleType(t.input[t.lastIndex:], ERROR, t.lastIndex))
	}
}

func (t *Tokenizer) getTokenFromBuffer(offset int, ignored []int) Token {
	for toRead := offset - len(t.buffer) + 1; toRead > 0; toRead-- {
		t.readTokenToBuffer()
	}

	token := t.buffer[offset]
	if len(ignored) > 0 && t.compareToken(token, ignored...) {
		res := NewTokenMultiTypes(token.Value, token.Types, token.Index)
		res.IsIgnored = true
		return res
	}
	return token
}

func (t *Tokenizer) compareToken(token Token, selectors ...int) bool {
	if len(selectors) < 1 {
		return false
	}

	for _, selector := range selectors {
		for _, tType := range token.Types {
			if tType == TokenType(selector) {
				return true
			}
		}
	}

	return false
}

func (t *Tokenizer) getTokenA(consume bool, ignored []int) TokenizerResult {
	offset := 0
	var token Token;
	if consume {
		for token = t.getTokenFromBuffer(0, ignored); token.IsIgnored; token = t.getTokenFromBuffer(0, ignored) {
			t.buffer = t.buffer[:len(t.buffer)]
		}
	} else {
		for token = t.getTokenFromBuffer(0, ignored); token.IsIgnored; token = t.getTokenFromBuffer(0, ignored) {
			offset++
		}
	}
	return NewTokenizerResult(token)
}

func (t *Tokenizer) getTokenB(consume bool, ignored []int, selector []int) TokenizerResult {
	count, index := 0, 0
	result := make([]Token, 0)

	for end := true; end; end = index < len(selector) {
		token := t.getTokenFromBuffer(count, ignored)
		count++
		if !token.IsIgnored && t.compareToken(token, selector[index]) {
			result = append(result, token)
			index++
		} else if !token.IsIgnored {
			return NotFoundTokenizerResult()
		}
	}

	if !consume {
		return FoundTokenizerResult()
	}

	t.buffer = t.buffer[count:]

	return NewTokenizerResult(result...);
}

func (t *Tokenizer) Next(ignored []int, selector ...int) *TokenizerResult {
	var res TokenizerResult
	if len(selector) > 0 {
		res = t.getTokenB(true, ignored, selector)
	} else {
		res = t.getTokenA(true, ignored)
	}
	return &res
}

func (t *Tokenizer) Test(ignored []int, selector ...int) *TokenizerResult {
	var res TokenizerResult
	if len(selector) > 0 {
		res = t.getTokenB(false, ignored, selector)
	} else {
		res = t.getTokenA(false, ignored)
	}
	return &res
}

func (t *Tokenizer) GetLineNumber(index int64) int {
	var pos int64 = -1;
	lineNumber := 1;
	for pos++; pos < index; pos++ {
		//todo
		//code := t.input[pos]
		//if code == "\r" || code == "\n" {
		//	lineNumber++
		//}
	}

	return lineNumber
}

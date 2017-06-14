package parser

import (
	"testing"
)

func TestNewTokenizer(t *testing.T) {
	input := "{{var a = 5}}{{a -> toJSON}}"
	tokenizer := NewTokenizer(input, "")

	assertToken(t, "{{", tokenizer)
	assertToken(t, "var", tokenizer)
	assertToken(t, " ", tokenizer)
	assertToken(t, "a", tokenizer)
	assertToken(t, " ", tokenizer)
	assertToken(t, "=", tokenizer)
	assertToken(t, " ", tokenizer)
	assertToken(t, "5", tokenizer)
	assertToken(t, "}}", tokenizer)
	assertToken(t, "{{", tokenizer)
	assertToken(t, "a", tokenizer)
	assertToken(t, " ", tokenizer)
	assertToken(t, "->", tokenizer)
	assertToken(t, " ", tokenizer)
	assertToken(t, "toJSON", tokenizer)
	assertToken(t, "}}", tokenizer)
	assertFailToFound(t, "}}", tokenizer)


}
func assertFailToFound(t *testing.T, str string, tokenizer *Tokenizer) {
	_, found := getNext(tokenizer)
	if found {
		t.Fatal("The token is found, but this should not be")
	}
}

func assertToken(t *testing.T, expectedToken string, tokenizer *Tokenizer) {
	token, found := getNext(tokenizer)
	if !found {
		t.Fatalf("Failed to found token")
	}

	if token.Value != expectedToken {
		t.Fatalf("Expected token is '%s', but found is '%s'", expectedToken, token.Value)
	}
}

func getNext(tokenizer *Tokenizer) (*Token, bool) {
	ignored := make([]int, 0)
	for i := 1; i <= 68; i++ {
		if tokenizer.Test(ignored, i).IsFound {
			token := tokenizer.Next(ignored, i).First()
			return &token, true
		}
	}
	return nil, false
}

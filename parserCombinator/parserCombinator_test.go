package parserCombinator

import (
	"testing"
)

func parseParen (input string) (ParseNode, string) {
	if(input[0] == '(') {
		return "PAREN", input[1:]
	} else {
		return nil, input
	}
}

var parseComma = func(input string) (ParseNode, string) {
	if(input[0] == ',') {
		return "COMMA", input[1:]
	} else {
		return nil, input
	}
}

var parseQuote = func(input string) (ParseNode, string) {
	if(input[0] == '"') {
		return "QUOTE", input[1:]
	} else {
		return nil, input
	}
}

func testParser(t *testing.T) {
	passThrough := func(nodes []ParseNode) ParseNode {
		return nodes
	}

	testParser := And(passThrough, parseQuote, parseParen)

	t.Log(testParser("\"("))
	t.Log(testParser("("))
	t.Log(testParser("other"))
}

func TestOr(t *testing.T) {
	passThrough := func(nodes []ParseNode) ParseNode {
		return nodes
	}

	testParser := Or(passThrough, parseQuote, parseParen)

	t.Log(testParser("\"("))
	t.Log(testParser("("))
	t.Log(testParser("other"))
}

func TestMaybe(t *testing.T) {
	passThrough := func(nodes []ParseNode) ParseNode {
		return nodes
	}

	testParser := Maybe(passThrough, parseParen)

	t.Log(testParser("\"("))
	t.Log(testParser("("))
	t.Log(testParser("other"))
}

func TestListOf(t *testing.T) {
	passThrough := func(nodes []ParseNode) ParseNode {
		return nodes
	}

	testParser := ListOf(passThrough, parseParen, parseComma)

	t.Log(testParser("\"("))
	t.Log(testParser("(,(,(stuff"))
	t.Log(testParser("(,(,(,stuff"))
	t.Log(testParser("other"))
}

func TestEatWhiteSpace(t *testing.T) {
	testStr := "   \t \r\n other"
	t.Log(testStr)
	t.Log(EatWhiteSpace(testStr))
}

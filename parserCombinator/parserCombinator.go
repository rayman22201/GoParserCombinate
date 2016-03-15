// Basically my own homegrown version of a Parser Combinator library like Parsec,
// With Blackjack....and Hookers :-P
// @see https://godoc.org/github.com/prataprc/goparsec
// @see https://en.wikipedia.org/wiki/Parser_combinator
// @see http://www.codecommit.com/blog/scala/the-magic-behind-parser-combinators
//
// The Basic idea is that you build up a complicated Parser by composing lots of tiny parsers.
// It's a very powerful tool for building small DSL's and Parsing type tasks.
// The code is much more readable than something from a Parser generator like YACC or BISON
//
package parserCombinator

import (
	"strings"
	"regexp"
)

// Empty interface can represent anything at all.
// @see http://blog.golang.org/laws-of-reflection
type ParseNode interface{}

type Nodify func(nodes...ParseNode) ParseNode

// parser interface
type Parser func(input string) (ParseNode, string)

// combinator interface
// The basic rule for Parser Combinators is that they take a function that processes the
// results of all the parsers first, then some number of parsers to combine in some way.
type Combinator func(resuls Nodify, parsers...Parser) Parser

// Combinators
func And(processResults Nodify, parsers...Parser) Parser {
	return func(input string) (ParseNode, string) {
		var resultNodes []ParseNode
		var output string

		output = input
		for _, curParser := range parsers {
			var result ParseNode
			result, output = curParser(output)
			// If we didn't match even one input, we failed.
			if result == nil {
				return nil, input
			} else {
				resultNodes = append(resultNodes, result)
				input = output
			}
		}
		if len(resultNodes) != len(parsers) {
			return nil, input
		}
		return processResults(resultNodes...), output
	};
}

// Note, this is a theoretically incorrect implementation of the Or combinator.
// This will greedily match on the first parser that succeeds,
// but the Theory states that an Or Combinator needs to be Exclusive Or.
// In other words, it must match EXACTLY one parser, no more, no less.
// This function, matches the first one it finds, another parser may also match further down
// the list, but we ignore that for practicality.
func Or(processResults Nodify, parsers...Parser) Parser {
	return func(input string) (ParseNode, string) {
		var output string

		output = input
		for _, curParser := range parsers {
			// If we encountered the end of the input before matching all options, we failed.
			if len(output) == 0 {
				return nil, input
			}

			var result ParseNode
			result, output = curParser(output)
			// If we match even one, then we succeeded
			if result != nil {
				return processResults(result), output
			}
		}

		// If we didn't match any of the results, we failed.
		return nil, input
	}
}

func Maybe(processSuccess Nodify, parse Parser) Parser {
	return func(input string) (ParseNode, string) {
		result, output := parse(input)
		if result == nil {
			return "", output
		}
		return processSuccess(result), output
	}
}

func ListOf(processResults Nodify, parse Parser, delimeter Parser) Parser {
	return func(input string) (ParseNode, string) {
		var results []ParseNode
		var output string
		var delimCount int
		output = input

		for {
			if len(output) == 0 {
				break
			}
			var result, delim ParseNode
			result, output = parse(output)

			if result != nil {
				results = append(results, result)
			} else {
				break
			}

			if len(output) == 0 {
				break
			}
			delim, output = delimeter(output)
			if delim == nil {
				break
			} else {
				delimCount++
			}
		}

		if len(results) > 0 && delimCount == (len(results) - 1) {
			return processResults(results...), output
		} else {
			return nil, input
		}
	}
}

func CharParser(search byte, terminal string) Parser {
	return func(input string) (ParseNode, string) {
		if(len(input) > 0 && input[0] == search) {
			return terminal, input[1:]
		} else {
			return nil, input
		}
	}
}

func RegexpParser(matcher *regexp.Regexp) Parser {
	return func(input string) (ParseNode, string) {
		matchStr := matcher.FindString(input)
		if matchStr == "" {
			return nil, input
		}
		return matchStr, strings.Replace(input, matchStr, "", 1)
	}
}

func PassAllNodes(nodes...ParseNode) ParseNode {
	return nodes
}

func PassNthNode(n int) Nodify {
	return func(nodes...ParseNode) ParseNode { return nodes[n-1] }
}

func EatWhiteSpace(input string) string {
	var spaceCode = make(map[byte]byte)
	spaceCode['\t'] = 1
	spaceCode['\n'] = 1
	spaceCode['\v'] = 1
	spaceCode['\f'] = 1
	spaceCode['\r'] = 1
	spaceCode[' '] = 1
	for {
		if _, ok := spaceCode[input[0]]; ok {
			input = input[1:]
		} else {
			break;
		}
	}
	return input
}

func EatEmptieNodes(nodes...ParseNode) ParseNode {
	var cleanedNodes []ParseNode
	for _, node := range nodes {
		switch t := node.(type) {
		case string:
			if t != "" {
				cleanedNodes = append(cleanedNodes, node)
			}
		default:
			cleanedNodes = append(cleanedNodes, node)
		}
	}
	return cleanedNodes
}
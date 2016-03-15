# Basically my own homegrown version of a Parser Combinator library like Parsec,

@see https://godoc.org/github.com/prataprc/goparsec
@see https://en.wikipedia.org/wiki/Parser_combinator

The Basic idea is that you build up a complicated Parser by composing lots of tiny parsers.
It's a very powerful tool for building small DSL's and Parsing type tasks.
The code is much more readable than something from a Parser generator like YACC or BISON

Originally created for a project done for Apartments 247, but never ended up being used.

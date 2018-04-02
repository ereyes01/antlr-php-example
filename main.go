package main

import (
	"flag"
	"fmt"
	"log"

	"antlr-php/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var filename string

type PHPScanner struct {
	*parser.BasePhpParserListener
}

func (s *PHPScanner) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("RULE:", ctx.GetText())
}

func init() {
	flag.StringVar(&filename, "file", "", "file to parse")
}

func main() {
	flag.Parse()
	if filename == "" {
		log.Fatalln("you must supply a filename")
	}

	input, err := antlr.NewFileStream(filename)
	if err != nil {
		log.Fatal(err)
	}
	lexer := parser.NewPhpLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPhpParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.HtmlDocument()
	antlr.ParseTreeWalkerDefault.Walk(new(PHPScanner), tree)
}

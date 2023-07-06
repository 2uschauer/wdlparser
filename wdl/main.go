package main

import (
	"fmt"

	"wdl/parser"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseWdlV1ParserListener
	p *parser.WdlV1Parser
	t antlr.Tree
}

func NewTreeShapeListener(p *parser.WdlV1Parser, t antlr.Tree) *TreeShapeListener {
	return &TreeShapeListener{
		p: p,
		t: t,
	}
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Printf("==> %s 《 %s 》\n", ruleName, ctx.GetText())
}

func (l *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	printLevelPrefix(ctx)
	i := ctx.GetRuleIndex()
	ruleName := l.p.RuleNames[i]
	fmt.Println("<==", ruleName)
}

func printLevelPrefix(ctx antlr.ParserRuleContext) {
	level := 0

	t := ctx.GetParent()
	for t != nil {
		level++
		t = t.GetParent()
	}

	for i := 0; i < level; i++ {
		fmt.Printf("\t")
	}
}

func main() {
	input, _ := antlr.NewFileStream("/Users/bytedance/Downloads/wdl/no.wdl")
	lexer := parser.NewWdlV1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewWdlV1Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Document()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(p, tree), tree)
}

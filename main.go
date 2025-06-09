package main

import (
	"fmt"
	"ytsejam/ast"
)

func main() {
	// Build the following program:
	// a := 5 + 3; b := (print(a, a-1), 10*a); print(b)

	prog := ast.NewCompoundStm(
		ast.NewAssignStm("a", ast.NewOpExp(
			ast.NewNumExp(5),
			ast.Plus,
			ast.NewNumExp(3),
		)),
		ast.NewCompoundStm(
			ast.NewAssignStm("b", ast.NewEseqExp(
				ast.NewPrintStm(ast.NewPairExpList(
					ast.NewIdExp("a"),
					ast.NewLastExpList(ast.NewOpExp(
						ast.NewIdExp("a"),
						ast.Minus,
						ast.NewNumExp(1),
					)),
				)),
				ast.NewOpExp(
					ast.NewNumExp(10),
					ast.Times,
					ast.NewIdExp("a"),
				),
			)),
			ast.NewPrintStm(ast.NewLastExpList(
				ast.NewIdExp("b"),
			)),
		),
	)

	// For now just print that we have built it
	fmt.Printf("%#v\n", prog)
}

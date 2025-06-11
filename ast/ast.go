package main

import "fmt"

type Binop int

const (
	A_plus Binop = iota
	A_minus
	A_times
	A_div
)

// // statements area \\\\\
type StmKind int

const (
	A_compoundStm StmKind = iota
	A_assignStm
	A_printStm
)

// /expressions area\\\\
type ExpKind int

const (
	A_idExp ExpKind = iota
	A_numExp
	A_opExp
	A_eseqExp
)

// Expression List Kind\\
type ExpListKind int

const (
	A_pairExpList ExpListKind = iota
	A_lastExpList
)

// Statement
type Stm struct {
	Kind     StmKind
	Compound *CompoundStm
	Assign   *AssignStm
	Print    *PrintStm
}

type CompoundStm struct {
	Stm1 *Stm
	Stm2 *Stm
}

type AssignStm struct {
	Id  string
	Exp *Exp
}

type PrintStm struct {
	Exps *ExpList
}

// Expression
type Exp struct {
	Kind ExpKind
	Id   string
	Num  int
	Op   *OpExp
	Eseq *EseqExp
}

type OpExp struct {
	Left  *Exp
	Oper  Binop
	Right *Exp
}

type EseqExp struct {
	Stm *Stm
	Exp *Exp
}

// Expression List
type ExpList struct {
	Kind ExpListKind
	Pair *PairExpList
	Last *Exp
}

type PairExpList struct {
	Head *Exp
	Tail *ExpList
}

// Constructors

// a statement that combines two statements
func Compound(stm1, stm2 *Stm) *Stm {
	return &Stm{Kind: A_compoundStm, Compound: &CompoundStm{Stm1: stm1, Stm2: stm2}}
}

// a statement that assigns an expression to an identifier
func Assign(id string, exp *Exp) *Stm {
	return &Stm{Kind: A_assignStm, Assign: &AssignStm{Id: id, Exp: exp}}
}

// a statement that prints expressions
func Print(exps *ExpList) *Stm {
	return &Stm{Kind: A_printStm, Print: &PrintStm{Exps: exps}}
}

// find and identify an identifier expression
func IdExp(id string) *Exp {
	return &Exp{Kind: A_idExp, Id: id}
}

// find and identify a number expression
func NumExp(n int) *Exp {
	return &Exp{Kind: A_numExp, Num: n}
}

// find and identify an operation expression (left item | operator | right item)
func OpExpFn(left *Exp, oper Binop, right *Exp) *Exp {
	return &Exp{Kind: A_opExp, Op: &OpExp{Left: left, Oper: oper, Right: right}}
}

// find and indentify an (statement-expression) sequence
func EseqExpFn(stm *Stm, exp *Exp) *Exp {
	return &Exp{Kind: A_eseqExp, Eseq: &EseqExp{Stm: stm, Exp: exp}}
}

// find and identify  the first expression and the next
func PairExpListFn(head *Exp, tail *ExpList) *ExpList {
	return &ExpList{Kind: A_pairExpList, Pair: &PairExpList{Head: head, Tail: tail}}
}

// find and identify the last expression on the list
func LastExpListFn(last *Exp) *ExpList {
	return &ExpList{Kind: A_lastExpList, Last: last}
}

// Example usage
func main() {
	// Example: print(5 + 3, 10)
	exp1 := OpExpFn(NumExp(5), A_plus, NumExp(3))
	exp2 := NumExp(10)
	explist := PairExpListFn(exp1, LastExpListFn(exp2))
	printStm := Print(explist)

	fmt.Printf("%+v\n", printStm)
}

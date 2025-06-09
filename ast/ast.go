package ast

// BinOp defines binary operators
type BinOp int

const (
	Plus BinOp = iota
	Minus
	Times
	Div
)

// === Statement (Stm) ===

type Stm interface{}

type CompoundStm struct {
	Stm1 Stm
	Stm2 Stm
}

type AssignStm struct {
	Id  string
	Exp Exp
}

type PrintStm struct {
	Exps ExpList
}

// === Expression (Exp) ===

type Exp interface{}

type IdExp struct {
	Id string
}

type NumExp struct {
	Num int
}

type OpExp struct {
	Left  Exp
	Oper  BinOp
	Right Exp
}

type EseqExp struct {
	Stm Stm
	Exp Exp
}

// === Expression List (ExpList) ===

type ExpList interface{}

type PairExpList struct {
	Head Exp
	Tail ExpList
}

type LastExpList struct {
	Last Exp
}

// === Factory Functions ===

// Statements
func NewCompoundStm(stm1, stm2 Stm) Stm {
	return CompoundStm{Stm1: stm1, Stm2: stm2}
}

func NewAssignStm(id string, exp Exp) Stm {
	return AssignStm{Id: id, Exp: exp}
}

func NewPrintStm(exps ExpList) Stm {
	return PrintStm{Exps: exps}
}

// Expressions
func NewIdExp(id string) Exp {
	return IdExp{Id: id}
}

func NewNumExp(num int) Exp {
	return NumExp{Num: num}
}

func NewOpExp(left Exp, oper BinOp, right Exp) Exp {
	return OpExp{Left: left, Oper: oper, Right: right}
}

func NewEseqExp(stm Stm, exp Exp) Exp {
	return EseqExp{Stm: stm, Exp: exp}
}

// Expression List
func NewPairExpList(head Exp, tail ExpList) ExpList {
	return PairExpList{Head: head, Tail: tail}
}

func NewLastExpList(last Exp) ExpList {
	return LastExpList{Last: last}
}

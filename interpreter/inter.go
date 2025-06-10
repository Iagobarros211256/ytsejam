package inter

import (
	"fmt"
)

type Env map[string]int

func InterpretStm(env Env, stm Stm) {
	switch s := stm.(type) {
	case CompoundStm:
		InterpretStm(env, s.Stm1)
		InterpretStm(env, s.Stm2)
	case AssignStm:
		val := InterpretExp(env, s.Exp)
		env[s.Id] = val
	case PrintStm:
		printExpList(env, s.Exps)
		fmt.Println() // newline after print
	}
}

func InterpretExp(env Env, exp Exp) int {
	switch e := exp.(type) {
	case IdExp:
		return env[e.Id]
	case NumExp:
		return e.Num
	case OpExp:
		left := InterpretExp(env, e.Left)
		right := InterpretExp(env, e.Right)
		switch e.Oper {
		case Plus:
			return left + right
		case Minus:
			return left - right
		case Times:
			return left * right
		case Div:
			return left / right
		default:
			panic("unknown operator")
		}
	case EseqExp:
		InterpretStm(env, e.Stm)
		return InterpretExp(env, e.Exp)
	default:
		panic("unknown expression")
	}
}

func printExpList(env Env, exps ExpList) {
	switch l := exps.(type) {
	case PairExpList:
		fmt.Printf("%d ", InterpretExp(env, l.Head))
		printExpList(env, l.Tail)
	case LastExpList:
		fmt.Printf("%d", InterpretExp(env, l.Last))
	default:
		panic("unknown exp list")
	}
}

func main() {
	env := make(Env)

	// Programa: a := 5 + 3; print(a, a - 1)
	prog := NewCompoundStm(
		NewAssignStm("a", NewOpExp(NewNumExp(5), Plus, NewNumExp(3))),
		NewPrintStm(
			NewPairExpList(
				NewIdExp("a"),
				NewLastExpList(
					NewOpExp(NewIdExp("a"), Minus, NewNumExp(1)),
				),
			),
		),
	)

	InterpretStm(env, prog)
}

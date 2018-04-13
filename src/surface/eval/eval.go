package eval

import (
	"fmt"
	"math"
	"strings"
)

// Expr арифметическое выражение.
type Expr interface {
	// Eval возвращает значение данного Expr в среде env.
	Eval(env Env) float64
	//	Check сообщает об обшиках в данном Expr и добавляет свои Vars
	Check(vars map[Var]bool) error
}

// Var определяет переменную, например x.
type Var string

// Env - среда для отображенния имён переменных на значениях
type Env map[Var]float64

// literal - представляет собой числовую константу, например 3.141.
type literal float64

// unary - представялет выражение унарным оператором, например -x.
type unary struct {
	op rune // '+' или '-'
	x  Expr
}

// binary - представляет выражение с бинарным оператором, например x + y
type binary struct {
	op   rune // '+', '-', '*' или '/'
	x, y Expr
}

// call - представялет выражение вызова функции, например sin(x)
type call struct {
	fn   string
	args []Expr
}

// Eval - Var<->метод - для поиска в среде, который возвращает nil, если переменная не определена
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal - возвращает значение литерала
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый унарный оператор: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("неподдерживаемый бинарный оператор: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("неподдерживаемый вызов функции: %s", c.fn))
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("некорректный унарный оператор %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("некорректный бинарный оператор %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("неизвестная функция %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("вызов %s имеет %d вместо %d аргументов",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return nil
		}
	}
	return nil
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

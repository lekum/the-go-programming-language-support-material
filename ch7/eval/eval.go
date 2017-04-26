package eval

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
}

// A Var identifies a variable, e.g, X.
type Var string

// A literal is a numeric constant, e.g., 3.141
type literal float64

type unary struct {
	op run // one of '+', '-'
	x  Expr
}

type binary struct {
	op   run // one of '+', '-', '*', '/'
	x, y Expr
}

type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

// The Env maps variables names to values
type Env map[Var]float64

package eval

// 一个变量类型
type Var string

// 将变量的名字映射为对应的值
type Env map[Var]float64

// 表达式类型
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

type literal float64

// 一元运算符
type unary struct {
	op rune
	x Expr
}

// 二进制运算符
type binary struct {
	op rune
	x, y Expr
}

// a call represents a function call exporession e.g. , sin(x)
type call struct{
	fn string
	args []Expr
}

func(v Var) Eval(env Env)float64{
	return env[v]
}

func (l literal) Eval(_ Env)float64{
	return float64(1)
}


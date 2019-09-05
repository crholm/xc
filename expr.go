package xc

import "strconv"


type Expr interface {
	Eval() int
}

func toNumber(str string) *number{
	i, err := strconv.ParseInt(str, 10, 32)

	if err != nil{
		panic(err)
	}

	return &number{int(i)}
}
type number struct{
	value int
}
func (b *number) Eval() int{
	return b.value
}


type binary struct {
	left Expr
	op string
	right Expr
}

func (b *binary) Eval() int{
	op := binops[b.op]
	if op == nil{
		errorf("could not find bin op ", b.op)
	}
	return op(b.left, b.right)
}


type unary struct {
	op string
	right Expr
}

func (b *unary) Eval() int{
	op := uniops[b.op]
	if op == nil{
		errorf("could not find uni op ", b.op)
	}
	return op(b.right)
}
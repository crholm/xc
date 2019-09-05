package xc

import (
	"crypto/rand"
	"math"
	"math/big"
)

var binops = map[string]func(left Expr, right Expr) int{}
var uniops = map[string]func(right Expr) int{}
func init(){

	binops["+"] = func(left Expr, right Expr) int {
		return left.Eval() + right.Eval()
	}
	binops["-"] = func(left Expr, right Expr) int {
		return left.Eval() - right.Eval()
	}
	binops["*"] = func(left Expr, right Expr) int {
		return left.Eval() * right.Eval()
	}
	binops["/"] = func(left Expr, right Expr) int {
		return left.Eval() / right.Eval()
	}
	binops["%"] = func(left Expr, right Expr) int {
		return left.Eval() % right.Eval()
	}
	binops["**"] = func(left Expr, right Expr) int {
		return int(math.Pow(float64(left.Eval()), float64(right.Eval())))
	}


	uniops["sqrt"] = func(right Expr) int {
		return int(math.Sqrt(float64(right.Eval())))
	}

	uniops["?"] = func(right Expr) int {
		max := right.Eval()
		i, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
		if err != nil{
			errorf("could not generate error %v", err)
		}
		return int(i.Int64())
	}

	uniops["!"] = func(right Expr) int {
		start := right.Eval()
		acc := start
		for i := 1; i < start; i++{
			acc *= start-i
		}
		return acc
	}


}

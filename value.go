package xc

import (
	"fmt"
	"math/big"
)

type valueType int

const (
	intType valueType = iota
	bigIntType
	bigFloatType
)


type Value interface {
	String() string
	Eval() Value
	toType(valueType) Value
}

type Int int64
func (i Int) toType(t valueType) Value {
	var v Value = i
	if t >= bigIntType {
		v = BigInt{
			big.NewInt(int64(i)),
		}
	}
	if t >= bigFloatType {
		v = BigFloat{
			new(big.Float).SetInt(v.(BigInt).Int),
		}
	}
	return v
}
func (i Int) String() string{
	return fmt.Sprintf("%d", i)
}
func (i Int) Eval() Value{
	return i
}


type BigInt struct {
	*big.Int
}
func (i BigInt) toType(t valueType) Value {
	if t == intType && i.IsInt64(){
		return Int(i.Int64())
	}
	if t == bigIntType {
		return i
	}
	if t == bigFloatType {
		v := BigFloat{
			new(big.Float),
		}
		v.SetInt(i.Int)
		return v
	}
	panic("could not down cast")
}
func (i BigInt) String() string{
	return i.String()
}
func (i BigInt) Eval() Value{
	return i
}


type BigFloat struct {
	*big.Float
}
func (i BigFloat) toType(t valueType) Value {
	if t == bigFloatType {
		return i
	}

	if t == bigIntType && i.IsInt() {
		ii, _:= i.Int(nil)
		return BigInt{ii}
	}

	if t == intType && i.IsInt() {
		ii, _:= i.Int(nil)
		if ii.IsInt64(){
			return Int(ii.Int64())
		}
	}
	panic("could not down cast")
}
func (i BigFloat) String() string{
	return i.String()
}
func (i BigFloat) Eval() Value{
	return i
}

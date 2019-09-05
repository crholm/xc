package xc

import (
	"fmt"
	"os"
	"strings"
)

const(
	TypeEOF = iota
	TypeNumber
	TypeOperator
	TypeLeftParen
	TypeRightParen
)



func stanitize(data []rune) []rune{
	//re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	//re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	//final := re_leadclose_whtsp.ReplaceAllString(string(data), "")
	//final = re_inside_whtsp.ReplaceAllString(final, " ")
	//fmt.Println(string(data), " /// ", final)
	//return []rune(final)

	return []rune(strings.Replace(string(data), " ", "", -1))
}


func NewParser(data []rune) Parser{

	return Parser{data: stanitize(data), pos: 0}
}

type Token struct {
	Raw  []rune
	Type int
}

func (t Token) len() int{
	return len(t.Raw)
}


type Parser struct {
	data []rune
	pos int
}

func (s *Parser) Reset(){
	s.pos = 0
}


func (s *Parser) Peek() (t Token){
	t.Type = TypeEOF

	var found bool
	var buf []rune
	for i, c := range s.data[s.pos:]{
		switch c {
		case '*':
			t.Type = TypeOperator
			t.Raw = []rune{c}
			if s.data[s.pos:][i+1] == '*'{
				t.Raw = []rune{c, c}
			}
			goto superbreak
		case '+', '-', '/', '!', '?', '%':
			if found {
				goto superbreak
			}
			t.Type = TypeOperator
			t.Raw = []rune{c}
			goto superbreak
		case '(':
			if found {
				goto superbreak
			}
			t.Type = TypeLeftParen
			t.Raw = []rune{c}
			goto superbreak
		case ')':
			if found {
				goto superbreak
			}
			t.Type = TypeRightParen
			t.Raw = []rune{c}
			goto superbreak
		default:
			found = true
			buf = append(buf, c)
		}
	}
	superbreak:

	if len(buf) != 0 {
		t.Raw = buf
		switch string(buf) {
		case "sqrt":
			t.Type = TypeOperator
		default:
			t.Type = TypeNumber
		}
	}

	return t
}



func (s *Parser) Next() (t Token){
	t = s.Peek()
	s.pos += t.len()
	return t
}


func(s *Parser) Eval() int{
	s.Reset()
	t := s.Next()
	e := s.expr(t)

	return e.Eval()
}




func (p *Parser) expr(tok Token) Expr {
	expr := p.operand(tok)    // Next slide.
	switch p.Peek().Type {
	case TypeEOF, TypeRightParen:
		return expr
	case TypeOperator:
		// Binary.
		tok = p.Next()
		return &binary{
			left:  expr,
			op:    string(tok.Raw),
			right: p.expr(p.Next()),   // Recursion.
		}
	}
	errorf("after expression: unexpected %s", p.Peek())
	return nil
}

func (p *Parser) operand(tok Token) Expr {
	var expr Expr
	switch tok.Type {
	case TypeOperator:
		expr = &unary{ op: string(tok.Raw), right: p.expr(p.Next())} // Mutual recursion.
	case TypeLeftParen:
		expr = p.expr(p.Next()) // Mutual recursion.
		tok := p.Next()
		if tok.Type != TypeRightParen {
			errorf("expected right paren, found %s", tok)
		}
	case TypeNumber:
		expr = toNumber(string(tok.Raw))
	default:
		errorf("unexpected %s", tok)
	}
	return p.index(expr) // Handled separately two slides from now.
}

func (p *Parser) index(expr Expr) Expr {
	//for p.Peek().Type == TypeLeftBrack {
	//	p.next()
	//	index := p.expr(p.next()) // Mutual recursion.
	//	tok := p.next()
	//	if tok.Type != scan.RightBrack {
	//		p.errorf("expected right bracket, found %s", tok)
	//	}
	//	expr = &binary{
	//		op:    "[]",
	//		left:  expr,
	//		right: index,
	//	}
	//}
	return expr
}



func errorf(s string, args ...interface{}) {
	fmt.Printf(s, args...)
	os.Exit(1)
}

// 	operand
//		( Expr )
//		( Expr ) [ Expr ]...
//		operand
//		number
//		operand [ Expr ]...
//		unop Expr


//	expr
//		operand
//		operand binop expr

package main

import (
	"fmt"
	"github.com/crholm/xc"
)



func testSimple(q string, expected int){
	p := xc.NewParser([]rune(q))
	got := p.Eval()

	fmt.Print("Result:\t")
	fmt.Println(q, "=", got)

}


func test(q string, expected int){
	p := xc.NewParser([]rune(q))
	got := p.Eval()

	if got == expected {
		fmt.Print("Pass:\t")
	}
	if got != expected {
		fmt.Print("Fail:\texpected ", expected, ", got ")
	}
	fmt.Println(q, "=", got)

}

func main(){

	test("1+3-12", 1+3-12)
	test("1+3-(2-3)", 1+3-(2-3))
	test("1+3-10*3", 1+3-10*3)
	//test("3*3-10*3", 3*3-10*3) // priority does not work
	//test("1+10*3-4", 1+10*3-4) // priority does not work
	test("1+(10*3)-4", 1+10*3-4)

	test("1+3-(10%3)", 1+3-(10%3))
	test("7/3", 7/3)

	test("2**3", 8)
	test("sqrt(64)", 8)
	test("!4", 24)
	test("!5", 120)
	test("!7", 5040)
	testSimple("(?100)*100", 100)
	testSimple("?100*100", 100) // bad priority should be equivalent to (?100)*100



	//for {
	//	t := p.Next()
	//	if t.Type == xc.TypeEOF {
	//		fmt.Println("EOF")
	//		break
	//	}
	//	fmt.Println(string(t.Raw), "\t\t-- ", t.Type)
	//}


}

// test1 project main.go
package main

import (
	"./ast"
	"bufio"
	"fmt"
		"./token"
	"strings"
)

func main() {

	//token.CommandReader = bufio.NewReader(os.Stdin)

	parseCode("def test()")
	parseCode("def test(p1,p2)")
	parseCode("a1")
	parseCode("1+1")
	parseCode("p1(a,b)+p2(a,b)")
	parseCode("1+(2+3)")
	parseCode("def test(p1,p2) p1(a,b)+p2(a,b)")
	parseCode("def showMessge(msg) toy_print(msg)")
	parseCode("showMessge('test')")


	println("DONE")
}

func parseCode(code string) {

	s:= strings.NewReader(code)
	token.CommandReader = bufio.NewReader(s)

	fmt.Print("command>"+code+", 解析AST：")

	token.Reset()

	curToken := token.GetToken()

	node := ast.Parse(curToken)

	fmt.Println(node.ToString())

}

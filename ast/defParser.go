package ast

import (
	"../token"
	"container/list"
)


/*
def test(x,y) x+y y
*/
func handleDEFStmt() *ASTNode {
	node := new(ASTNode)
	node.CreateDEFASTNode()

	token.GetToken()
	sig:=parseFunctionSignature()

	node.FunctionSignature=sig

	token.GetToken()
	node.FunctionBody=parseExpression()

	return node
}

func parseFunctionSignature() *FunctionSignature {

	sig:=new(FunctionSignature)

	if token.CurrentToken!=token.TOK_identifier{
		panic("无法找到函数名")
	}
	sig.Identifier=token.Identifier_value

	token.GetToken()
	if token.CurrentToken!=token.TOK_LB{
		panic("无法找到函数左括号")
	}
	plist:=list.New()
	for{
		token.GetToken()
		if token.CurrentToken==token.TOK_RB{		//碰到右括号
			break
		}
		if(token.CurrentToken==token.TOK_COMMA){	//eat 逗号
			token.GetToken()
		}

		if token.CurrentToken!=token.TOK_identifier{
			panic("无法找到参数")
		}
		plist.PushBack(token.Identifier_value)
	}

	pArray:=make([]Parameter, plist.Len())
	index:=0
	for v := plist.Front(); v != nil; v = v.Next() {
		p:=new(Parameter)
		p.identifier=v.Value.(string)

		pArray[index]= *p

		index++
	}
	sig.parameters=pArray

	return sig
}



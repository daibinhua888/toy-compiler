package ast

import (
	"../token"
	"container/list"
	)

/*
	a=100
	1+2
*/
func handleTopLevelExpression() *ASTNode {
	return parseExpression()
}

func parseNumericExpr() *ASTNode{
	node := new(ASTNode)
	node.CreateNumericExpressionASTNode()

	node.NumericValue=token.Numeric_value

	return node
}

func parseIdentifierExpr() *ASTNode  {
	node := new(ASTNode)
	node.CreateIdentifierExpressionASTNode()

	node.VariableName=token.Identifier_value

	token.GetToken()
	if token.CurrentToken!=token.TOK_LB{
		return node
	}
	plist:=list.New()
	for{
		token.GetToken()
		//fmt.Println(token.CurrentToken)
		if token.CurrentToken==token.TOK_RB{		//碰到右括号
			break
		}
		if(token.CurrentToken==token.TOK_COMMA){	//eat 逗号
			token.GetToken()
		}

		if(token.CurrentToken==token.TOK_identifier){
			tempParameter:=new(Parameter)
			tempParameter.identifier=token.Identifier_value
			tempParameter.parameterType="identifier"
			plist.PushBack(tempParameter)
		} else if(token.CurrentToken==token.TOK_STRING){
			tempParameter:=new(Parameter)
			tempParameter.identifier=token.String_value
			tempParameter.parameterType="string"
			plist.PushBack(tempParameter)
		}
	}

	pArray:=make([]Parameter, plist.Len())
	index:=0
	for v := plist.Front(); v != nil; v = v.Next() {
		pArray[index]=* v.Value.(* Parameter)
		index++
	}
	callSig:=new(FunctionCall)
	callSig.Identifier=node.VariableName
	callSig.parameters=pArray

	node.FunctionCall=callSig

	return node
}

func parsePrimary() *ASTNode {

	tokenId:=token.CurrentToken

	if tokenId==token.TOK_identifier{
		return parseIdentifierExpr()
	} else if tokenId==token.TOK_number{
		return parseNumericExpr()
	}else if tokenId==token.TOK_LB{
		return parseParenthesisExpression()
	}

	node := new(ASTNode)
	node.CreateEmptyASTNode()
	return node
}

func parseParenthesisExpression() *ASTNode {
	token.GetToken()//eat '('
	node:=parseExpression()

	return node
}

func parseExpression() *ASTNode {
	pNode := new(ASTNode)
	leftNode := parsePrimary()

	if token.GetCurrentChar() == "+" {
		pNode.Op = "+"
		pNode.NodeType="Operator"
		pNode.LHS = leftNode
		token.GetToken()
		token.GetToken()
		pNode.RHS = parsePrimary()
	} else {
		pNode = leftNode
	}

	return pNode
}
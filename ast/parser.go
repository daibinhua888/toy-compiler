package ast

import (
	"../token"
	)

func Parse(tk int) *ASTNode {
	switch tk {
	case token.TOK_def:
		return handleDEFStmt()
	default:
		return handleTopLevelExpression()
	}
}

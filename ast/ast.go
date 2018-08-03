// ast project ast.go
package ast

import "strconv"

type ASTNode struct {
	NodeType     string
	Op string
	NumericValue float64
	VariableName string
	LHS          *ASTNode
	RHS          *ASTNode
	FunctionSignature *FunctionSignature
	FunctionCall *FunctionCall
	FunctionBody *ASTNode
}

type Parameter struct {
	identifier string
}

type FunctionSignature struct {
	Identifier string
	parameters []Parameter
}

type FunctionCall struct {
	Identifier string
	parameters []Parameter
}

func (self *ASTNode) CreateEmptyASTNode() {
	self.NodeType = "empty"
}

func (self *ASTNode) CreateDEFASTNode() {
	self.NodeType = "def"
}

func (self *ASTNode) CreateExpressionASTNode() {
	self.NodeType = "expression"
}
func (self *ASTNode) CreateNumericExpressionASTNode() {
	self.NodeType = "numeric"
}

func (self *ASTNode) CreateIdentifierExpressionASTNode() {
	self.NodeType = "identifier"
}

func (self *ASTNode) ToString() string {
	str:=""

	str=self.NodeType+"-"+self.VariableName+"("+self.Op+")"
	if self.FunctionSignature!=nil{
		str+="-(SIG)"+self.FunctionSignature.Identifier
		if self.FunctionSignature.parameters!=nil{
			str+="("
			for index:=0;index< len(self.FunctionSignature.parameters);index++{
				str+=self.FunctionSignature.parameters[index].identifier+","
			}
			str+=")"
		}
	}

	if self.FunctionCall!=nil{
		str+="-(CALL)"+self.FunctionCall.Identifier
		if self.FunctionCall.parameters!=nil{
			str+="("
			for index:=0;index< len(self.FunctionCall.parameters);index++{
				str+=self.FunctionCall.parameters[index].identifier+","
			}
			str+=")"
		}
	}

	if self.LHS!=nil &&self.LHS.NodeType!="empty"{
		str+="(LHS: "+strconv.FormatFloat(self.LHS.NumericValue,'G',5,64)+")"
		str+=self.LHS.ToString()+"---"
	}
	if self.RHS!=nil &&self.RHS.NodeType!="empty"{
		str+="(RHS: "+strconv.FormatFloat(self.RHS.NumericValue,'G',5,64)+")"
		str+=self.RHS.ToString()+"---"
	}
	if(self.FunctionBody!=nil){
		str+="Body Expr: "+self.FunctionBody.ToString()
	}

	return str
}

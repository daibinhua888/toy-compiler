// ast project ast.go
package ast

import (
	"strconv"
	"strings"
)

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
	stringText string
	parameterType string
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
func (self *ASTNode) GenerateLLVMCode4DefineVariables() string {

	defineStringTemplate:=`@.{variableName} = private unnamed_addr constant [{stringLength} x i8] c"{stringContent}\0A\00"`

	codes:=""
	if self.FunctionCall!=nil{
		if self.FunctionCall.parameters!=nil{
			for i:=0;i<len(self.FunctionCall.parameters);i++{
				p:=self.FunctionCall.parameters[i]
				if p.parameterType=="string"{

					curCodeLine:=strings.Replace(defineStringTemplate, "{variableName}", strings.Trim(p.identifier, "'"), -1)
					length:=len(strings.Trim(p.identifier, "'"))+2
					curCodeLine=strings.Replace(curCodeLine, "{stringLength}", strconv.Itoa(length), -1)
					curCodeLine=strings.Replace(curCodeLine, "{stringContent}", strings.Trim(p.identifier, "'"), -1)

					codes+=curCodeLine+"\r\n"
				}
			}
		}
	}


	return codes
}
func (self *ASTNode) GenerateLLVMCode4DefineFunction() string{

	if self.FunctionSignature!=nil{
		defineFunctionTemplate:=`
define void @{functionName}(){
;
	{functionBody}
	ret void
}
`
		code:=strings.Replace(defineFunctionTemplate, "{functionName}", self.FunctionSignature.Identifier, -1)

		code=strings.Replace(code, "{functionBody}", generateBody(self.FunctionBody), -1)

		return code
	}

	return ""
}
func generateBody(node *ASTNode) string {

	if node.FunctionCall!=nil{
		if node.FunctionCall.Identifier=="toy_print"{
			//puts函数调用
			toy_print_template:=`
	%cast210 = getelementptr [6 x i8], [6 x i8]* @.test, i64 0, i64 0
	call i32 @puts(i8* %cast210)
`
			print_code:=toy_print_template
			return print_code
		}
	}

	return ""
}
func (self *ASTNode) GenerateLLVMCode4InMain() string{

	if self.FunctionCall!=nil{
		return "	call void @"+self.FunctionCall.Identifier+"()"
	}

	return ""
}

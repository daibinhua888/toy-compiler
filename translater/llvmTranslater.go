package translater

import (
	"../ast"
	"fmt"
	"strings"
)


var mainTemplate string=`

declare i32 @puts(i8* nocapture) nounwind

{codesInDefineVariables}
{codesInDefineFunction}

define i32 @main(){
;	
	{codesInMain}
	ret i32 0
}
`

var codesInDefineVariables string=""
var codesInDefineFunction string=""
var codesInMain string=""

func RunAST(astNode *ast.ASTNode)  {
	codesInDefineVariables+="\r\n"+astNode.GenerateLLVMCode4DefineVariables()
	codesInDefineFunction+="\r\n"+astNode.GenerateLLVMCode4DefineFunction()
	codesInMain+="\r\n"+astNode.GenerateLLVMCode4InMain()
}

func Dump() {
	llCodeContent:=strings.Replace(mainTemplate, "{codesInDefineVariables}", codesInDefineVariables, -1)
	llCodeContent=strings.Replace(llCodeContent, "{codesInDefineFunction}", codesInDefineFunction, -1)
	llCodeContent=strings.Replace(llCodeContent, "{codesInMain}", codesInMain, -1)
	fmt.Println(llCodeContent)
}
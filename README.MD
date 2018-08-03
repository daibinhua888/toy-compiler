用go写的玩具编译器，最终目标是后端套用LLVM生成本机字节码

#Code:
```go
parseCode("def test()")
parseCode("def test(p1,p2)")
parseCode("a1")
parseCode("1+1")
parseCode("p1(a,b)+p2(a,b)")
parseCode("1+(2+3)")
parseCode("def test(p1,p2) p1(a,b)+p2(a,b)")
```

#AST:
```go
command>def test(), 解析AST：def-()-(SIG)test()Body Expr: empty-()
command>def test(p1,p2), 解析AST：def-()-(SIG)test(p1,p2,)Body Expr: empty-()
command>a1, 解析AST：identifier-a1()
command>1+1, 解析AST：Operator-(+)(LHS: 1)numeric-()---(RHS: 1)numeric-()---
command>p1(a,b)+p2(a,b), 解析AST：Operator-(+)(LHS: 0)identifier-p1()-(CALL)p1(a,b,)---(RHS: 0)identifier-p2()-(CALL)p2(a,b,)---
command>1+(2+3), 解析AST：Operator-(+)(LHS: 1)numeric-()---(RHS: 0)Operator-(+)(LHS: 2)numeric-()---(RHS: 3)numeric-()------
command>def test(p1,p2) p1(a,b)+p2(a,b), 解析AST：def-()-(SIG)test(p1,p2,)Body Expr: Operator-(+)(LHS: 0)identifier-p1()-(CALL)p1(a,b,)---(RHS: 0)identifier-p2()-(CALL)p2(a,b,)---
```
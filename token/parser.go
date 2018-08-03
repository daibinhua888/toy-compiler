package token

import (
	"bufio"
	"strconv"
)

type token int

const (
	TOK_eof        = -1
	TOK_if         = -2
	TOK_then       = -3
	TOK_identifier = -4
	TOK_expression = -5
	TOK_def        = -6
	TOK_number     = -7
	TOK_LB=-8
	TOK_RB=-9
	TOK_COMMA=-10
	TOK_PLUS=-11
	TOK_MINUS=-12
	TOK_MULTIPLY=-13
	TOK_DIV=-14
)

var CommandReader *bufio.Reader = nil

var Identifier_value string
var Numeric_value float64

func getChar() (string, byte) {
	char_byte, _ := CommandReader.ReadByte()
	char := string(char_byte)
	return char, char_byte
}

var lastChar string = " "
var lastCharByte byte
var CurrentToken int

func getToken() int {
	for IsSpace(lastChar) {
		lastChar, lastCharByte = getChar()
	}

	if IsAlpha(lastChar) { //字母开头的：可能是变量名、def
		Identifier_value = lastChar

		for {
			lastChar, lastCharByte = getChar()
			if IsAlphaOrNumeric(lastChar) {
				Identifier_value += lastChar
			} else {
				break
			}
		}

		if Identifier_value == "def" {
			return TOK_def
		}

		return TOK_identifier
	}

	if IsNumeric(lastChar) {
		tmpNumber := lastChar
		for {
			lastChar, lastCharByte = getChar()
			if IsNumeric(lastChar) || lastChar == "." {
				tmpNumber += lastChar
			} else {
				break
			}
		}

		Numeric_value, _ = strconv.ParseFloat(tmpNumber, 32)
		return TOK_number
	}
	if lastChar=="("{
		lastChar, lastCharByte = getChar()
		return TOK_LB
	}
	if lastChar==")"{
		lastChar, lastCharByte = getChar()
		return TOK_RB
	}
	if lastChar==","{
		lastChar, lastCharByte = getChar()
		return TOK_COMMA
	}

	if lastChar=="+"{
		lastChar, lastCharByte = getChar()
		return TOK_PLUS
	}
	if lastChar=="-"{
		lastChar, lastCharByte = getChar()
		return TOK_MINUS
	}
	if lastChar=="*"{
		lastChar, lastCharByte = getChar()
		return TOK_MULTIPLY
	}
	if lastChar=="/"{
		lastChar, lastCharByte = getChar()
		return TOK_DIV
	}

	if lastCharByte == 13||lastCharByte == 10 {
		return TOK_eof
	}

	_, cur_byte := lastChar, lastCharByte
	lastChar, lastCharByte = getChar()
	return int(cur_byte)
}
func GetToken() int {
	CurrentToken=getToken()
	return CurrentToken
}

func DisplayInfo(){
	println("============")
	println("CurrentToken: "+strconv.Itoa(CurrentToken))
	println("lastChar: "+lastChar)
	print("lastCharByte: ")
	println(lastCharByte)
	println("============")
}

func GetCurrentChar() string {
	return lastChar
}

func Reset() {
	lastChar=" "
	lastCharByte=0
	CurrentToken=0
}
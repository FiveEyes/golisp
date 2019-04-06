package core

import (
	"fmt"
	"strconv"
)

func init() {
	//RegisterNativeFunction("parse", Parse, BasicEnv)
	fmt.Println("parser.init done")
}

func parseNum(s string, i int) (XLObj, int) {
	j := i
	n := len(s)
	f := false
	var v XLObj
	for ; j < n; j++ {
		if '0' <= s[j] && s[j] <= '9' {
			continue
		} else if s[j] == '.' {
			f = true
			continue
		} else {
			break
		}
	}
	if !f {
		x, _ := strconv.Atoi(s[i:j])
		v = NewXLInt(x)
	} else {
		x, _ := strconv.ParseFloat(s[i:j], 64)
		v = NewXLFloat(x)
	}
	return v, j
}

func parseString(s string, i int) (*XLString, int) {
	n := len(s)
	j := i
	for ; j < n; j++ {
		if s[j] == '"' {
			break
		}
	}
	return NewXLString(s[i+1:j]), j+1
}

func parseSymbol(s string, i int) (*XLSymbol, int) {
	n := len(s)
	j := i
	for ; j < n; j++ {
		//if 'A' <= s[j] && s[j] <= 'Z' || 'a' <= s[j] && s[j] <= 'z' || '0' <= s[j] && s[j] <= '9' {
		if s[j] == ' ' || s[j] == ')' || s[j] == '(' {
			break
		}
		continue
	}
	return NewXLSymbol(s[i:j]), j
}

func String2tokens(s string) []XLObj {
	var l []XLObj
	i := 0
	n := len(s)
	left := NewXLSymbol("(")
	right := NewXLSymbol(")")
	for ; i < n; {
		//fmt.Println(i, s[i:i+1])
		var x XLObj
		if s[i] == ' ' {
			i++
			continue
		} else if s[i] == '(' {
			x = left
			i++
		} else if s[i] == ')' {
			x = right
			i++
		} else if '0' <= s[i] && s[i] <= '9' {
			x, i = parseNum(s, i)
		} else if s[i] == '"' {
			x, i = parseString(s, i)
		} else {
			x, i = parseSymbol(s, i)
		} 
		l = append(l, x)

	}
	return l
}

func list2XLObj(t []XLObj) XLObj {
	//fmt.Println("l2x", t)
	n := len(t)
	var p XLObj
	p = Nil
	for i:=n-1; i>=0; i-- {
		p = NewXLPair(t[i], p)
	}
	return p
}

func parseTokens(t []XLObj, i int) (XLObj, int) {
	if t[i].XLObjType() == DT_Symbol {
		if t[i].(*XLSymbol).Value == "(" {
			l := make([]XLObj, 0, 7)
			j := i+1
			n := len(t)
			for ; j < n; {
				if t[j].XLObjType() == DT_Symbol && t[j].(*XLSymbol).Value == ")" {
					break
				}
				var next XLObj
				next, j = parseTokens(t, j)
				l = append(l, next)
			}
			return list2XLObj(l), j+1
		} else {
			return t[i], i+1
		}
	} else {
		return t[i], i+1
	}
}
func Parse(s string) (XLObj, bool) {
	tokens := String2tokens(s)
	//fmt.Println(tokens)
	x, _ := parseTokens(tokens, 0)
	return x, true
}

 
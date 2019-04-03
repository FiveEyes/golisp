package main

import (
	"fmt"
	"xlisp/pkg/xlcore"
)

func evalString(s string, env core.XLEnv) (core.XLObj, bool) {
	fmt.Println("in:", s)
	p, _ := core.Parse(s)
	//fmt.Println("parse:", core.PrettyPrint(p))
	ret, ok := core.ExpEval(p, env)
	if !ok {
		fmt.Println("out:", "error")
	} else {
		if ret.XLObjType() != core.DT_Function {
			fmt.Println("out:", core.PrettyPrint(ret))
		} else {
			fmt.Println("out:", "function")
		}
	}
	return ret, ok
}

func main() {
	/*
	fmt.Println("hello")
	x := core.XLInt {Value: 5}
	z, _ := core.Eval(x, core.NewEnv())
	fmt.Println(z)
	
	env := core.NewEnv()
	core.EnvPut(env, "a", x)
	a := core.XLSymbol {Value : "a"}
	fmt.Println(core.Eval(a, env))	

	core.EnvPut(env, "+", core.XLNativeFunction{Func : core.Add})
	addsym := core.XLSymbol {Value : "+"}
	p1 := core.XLPair{Fst: x, Snd: core.XLNil{}}
	p2 := core.XLPair{Fst: x, Snd: p1}
	p3 := core.XLPair{Fst: addsym, Snd: p2}
	fmt.Println(core.Eval(p3, env))
	p4, _ := core.Parse("(+ 1 2.5)")
	fmt.Println(p4)
	fmt.Println(core.Eval(p4, core.BasicEnv))
	
	p5, _ := core.Parse("(eval (quote (+ 1 2.5)))")
	fmt.Println(p5)
	fmt.Println(core.ExpEval(p5, core.BasicEnv))
	p6, _ := core.Parse("(quote (+ 1 2.5))")
	fmt.Println(core.ExpEval(p6, core.BasicEnv))
	*/
	
	env := core.BasicEnv
	evalString("(define x 1.5)", env)
	evalString("x", env)
	evalString("(lambda (y) (+ x y))", env)
	evalString("((lambda (y) (+ x y)) 1)", env)
	evalString("(quote (define add1 (lambda (y) (+ x y))))", env)
	evalString("(eval (quote (define add1 (lambda (y) (+ x y)))))", env)
	evalString("(add1 x)", env)
	evalString("(define add3 (lambda (x y z) (+ x y z)))", env)
	evalString("(add3 1 2 3)", env)
	
	evalString("(define fib (lambda (n) (if (eq? n 0) 1 (if (eq? n 1) 1 (+ (fib (- n 1)) (fib (- n 2)))))))", env)
	evalString("(fib 5)", env)

	evalString("(define fib1 (lambda (n) (if (eq? n 0) (cons 1 nil) (if (eq? n 1) (quote (1 1)) (cons (+ (car (fib1 (- n 1))) (car (cdr (fib1 (- n 1))))) (fib1 (- n 1)))))))", env)
	evalString("(fib1 5)", env)
	
	evalString("(define fib2 (lambda (n) (if (eq? n 0) (cons 1 nil) (if (eq? n 1) (quote (1 1)) (let (tmp (fib2 (- n 1))) (cons (+ (car tmp) (car (cdr tmp))) tmp)))))", env)
	evalString("(fib2 10)", env)

	evalString("(define add4 (lambda (x) (lambda (y) (+ x y))))", env)
	evalString("(define add5 (add4 5))", env)
	evalString("(define add5 (add5 5))", env)
	
	
}	
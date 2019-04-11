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
	env := core.BasicEnv
	/*
	var ret core.XLObj
	
	fmt.Println("hello")
	x := core.NewXLInt(5)
	z, _ := core.ExpEval(x, env)
	fmt.Println(z)
	
	core.EnvPut(env, "a", x)
	a := core.NewXLSymbol("a")
	fmt.Println(core.ExpEval(a, env))	
	
	l := core.NewXLLazy(a, env)
	fmt.Println(&l, core.PrettyPrint(l))
	fmt.Println(core.ExpEval(l, env))	
	fmt.Println(core.PrettyPrint(l))

	
	//core.EnvPut(env, "+", core.XLNativeFunction{Func : core.Add})
	addsym := core.NewXLSymbol("+")
	p1 := core.NewXLPair(x, core.Nil)
	p2 := core.NewXLPair(x, p1)
	p3 := core.NewXLPair(addsym, p2)
	fmt.Println(core.Eval(p3, env))
	p4, _ := core.Parse("(- 2 1)")
	fmt.Println(core.PrettyPrint(p4))
	ret, _ = core.ExpEval(p4, core.BasicEnv)
	fmt.Println(core.PrettyPrint(ret))
	
	p5, _ := core.Parse("(eval (quote (+ 1 2.5)))")
	fmt.Println(p5)
	fmt.Println(core.ExpEval(p5, core.BasicEnv))
	p6, _ := core.Parse("(quote (+ 1 2.5))")
	fmt.Println(core.ExpEval(p6, core.BasicEnv))
	*/
	
	
	evalString("(def x 1.5)", env)
	evalString("x", env)
	evalString("(lam (y) (+ x y))", env)
	evalString("((lam (y) (+ x y)) 1)", env)
	evalString("(quote (def add1 (lam (y) (+ x y))))", env)
	evalString("(eval (quote (def add1 (lam (y) (+ x y)))))", env)
	evalString("(def x 2.5)", env)
	evalString("(add1 x)", env)
	evalString("(def add3 (lam (x y z) (+ x y z)))", env)
	evalString("(add3 1 2 3)", env)
	
	evalString("(def fib (lam (n) (if (eq? n 0) 1 (if (eq? n 1) 1 (+ (fib (- n 1)) (fib (- n 2)))))))", env)
	evalString("(fib 1)", env)
	evalString("(fib 2)", env)
	evalString("(fib 3)", env)
	evalString("(fib 4)", env)
	evalString("(fib 5)", env)
	
	evalString("(def fib1 (lam (n) (if (eq? n 0) (cons 1 nil) (if (eq? n 1) (quote (1 1)) (cons (+ (car (fib1 (- n 1))) (car (cdr (fib1 (- n 1))))) (fib1 (- n 1)))))))", env)
	evalString("(fib1 5)", env)
	
	evalString("(def fib2 (lam (n) (if (eq? n 0) (cons 1 nil) (if (eq? n 1) (quote (1 1)) (let (tmp (fib2 (- n 1))) (cons (+ (car tmp) (car (cdr tmp))) tmp)))))", env)
	evalString("(fib2 10)", env)

	evalString("(def add4 (lam (x) (lam (y) (+ x y))))", env)
	evalString("(def add5 (add4 5))", env)
	evalString("(def add5 (add5 5))", env)


	evalString("(car (lazy (quote (1 2 3)))", env)


	evalString("(def lx (lazy (+ 1 2)))", env)
	evalString("(def inffib (lam (a b) (cons a (lazy (inffib b (+ a b))))))", env)
	evalString("(def fibinf (inffib 1 1))", env)
	evalString("(cdr fibinf)", env)
	evalString("(cdr (cdr fibinf))", env)
	evalString("(cdr (cdr (cdr fibinf)))", env)
	evalString("(cdr (cdr (cdr (cdr fibinf))))", env)
	evalString("fibinf", env)
	evalString("(def take (lam (l n) (if (eq? n 0) nil (cons (car l) (take (cdr l) (- n 1))))))", env)
	evalString("(take (fib2 10) 3)", env)
	evalString("(take (inffib 1 1) 10)", env)

	
}
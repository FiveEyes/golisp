package core

import "fmt"


func init() {	
	//RegisterNativeFunction("+", Add, false, BasicEnv)
	//RegisterNativeFunction("-", Minus, false, BasicEnv)
	//RegisterNativeFunction("eq?", Equal, false, BasicEnv)
	
	RegisterNativeFunction("quote", Quote, true, BasicEnv)
	RegisterNativeFunction("lambda", Lambda, true, BasicEnv)
	RegisterNativeFunction("define", Define, true, BasicEnv)
	RegisterNativeFunction("let", Let, true, BasicEnv)
	RegisterNativeFunction("if", If, true, BasicEnv)
	RegisterNativeFunction("lazy", Lazy, true, BasicEnv)
	
	RegisterNativeFunction("cons", Cons, false, BasicEnv)
	RegisterNativeFunction("car", Car, false, BasicEnv)
	RegisterNativeFunction("cdr", Cdr, false, BasicEnv)
	
	RegisterNativeFunction("eval", Eval, false, BasicEnv)
	
	EnvPut(BasicEnv, "nil", Nil)
	
	fmt.Println("basicNF.init done")
}



func Quote(objLazy XLObj, env XLEnv) (XLObj, bool) {
	return (objLazy.(*XLLazy).Value).(*XLPair).Fst, true
}

func Car(obj XLObj, env XLEnv) (XLObj, bool) {
	//fmt.Println("car", PrettyPrint(obj))
	return GetValue(obj.(*XLPair).Fst).(*XLPair).Fst, true
}

func Cdr(obj XLObj, env XLEnv) (XLObj, bool) {
	//fmt.Println("cdr", PrettyPrint(obj))
	return GetValue(obj.(*XLPair).Fst).(*XLPair).Snd, true
}

func Cons(obj XLObj, env XLEnv) (XLObj, bool) {
	return NewXLPair(PairCar(obj), PairCar(PairCdr(obj))), true
}

func Define(objLazy XLObj, env XLEnv) (XLObj, bool) {
	obj := objLazy.(*XLLazy).Value
	name := obj.(*XLPair).Fst
	exp := obj.(*XLPair).Snd.(*XLPair).Fst
	if name.XLObjType() == DT_Symbol {
		ret, ok := ExpEval(exp, env)
		if ok {
			EnvPut(env, name.(*XLSymbol).Value, ret)
			return ret, ok
		}
	}
	
	return Nil, false
	
}

func Let(objLazy XLObj, env XLEnv) (XLObj, bool) {
	obj := objLazy.(*XLLazy).Value
	nenv := PushNewEnv(env)
	var p XLObj
	for p = obj; p.(*XLPair).Snd.XLObjType() != DT_Nil; p = p.(*XLPair).Snd {
		name := p.(*XLPair).Fst.(*XLPair).Fst
		exp := p.(*XLPair).Fst.(*XLPair).Snd.(*XLPair).Fst
		if name.XLObjType() == DT_Symbol {
			ret, ok := ExpEval(exp, nenv)
			//fmt.Println(name, PrettyPrint(exp), PrettyPrint(ret))

			if ok {
				EnvPut(nenv, name.(*XLSymbol).Value, ret)
			} else {
				return Nil, false

			}
		} else {
			return Nil, false

		}
	}
	//fmt.Println(PrettyPrint(p.(XLPair).Fst))
	return ExpEval(p.(*XLPair).Fst, nenv)
	
}

func symlist2strings(obj XLObj) []string {
	n := ListLength(obj)
	s := make([]string, n)
	i := 0
	for p := obj; p.XLObjType() != DT_Nil; p = p.(*XLPair).Snd {
		s[i] = p.(*XLPair).Fst.(*XLSymbol).Value
		i += 1
	}
	return s
}

func Lambda(objLazy XLObj, env XLEnv) (XLObj, bool) {
	obj := objLazy.(*XLLazy).Value
	f := NewXLFunction(symlist2strings(obj.(*XLPair).Fst), obj.(*XLPair).Snd.(*XLPair).Fst, env, false)
	return f, true
}

func If(objLazy XLObj, env XLEnv) (XLObj, bool) {
	obj := objLazy.(*XLLazy).Value
	c := PairCar(obj)
	fst := PairCar(PairCdr(obj))
	snd := PairCar(PairCdr(PairCdr(obj)))
	t, ok := ExpEval(c, env)
	if !ok {
		return Nil, false
	}
	if t.(*XLInt).Value == 1 {
		return ExpEval(fst, env)
	} else {
		return ExpEval(snd, env)
	}
}

func Lazy(objLazy XLObj, env XLEnv) (XLObj, bool) {
	obj := objLazy.(*XLLazy)
	fst := PairCar(obj.Value)
	return NewXLLazy(fst, obj.Env), true
}


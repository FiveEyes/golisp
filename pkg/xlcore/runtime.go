package core

import (
	"fmt"
)

func init() {
	RegisterNativeFunction("eval", Eval, false, BasicEnv)
	fmt.Println("runtime.init done")
}

//XLEnv List XLMap
type XLEnv *XLPair

var BasicEnv XLEnv = NewXLEnv()

func NewXLEnv() XLEnv {
	return XLEnv(NewXLPair(NewXLMap(), Nil))
}

func EnvGet(env XLEnv, s string) (XLObj, bool) {
	m := env.Fst.(*XLMap)
	v, ok := MapGet(m, s)
	if ok {
		return v, true
	} else if env.Snd.XLObjType() == DT_Nil {
		return Nil, false
	} else {
		return EnvGet(XLEnv(env.Snd.(*XLPair)), s)
	}
}

func EnvPut(env XLEnv, s string, v XLObj) {
	MapPut(env.Fst.(*XLMap), s, v)
}

func PushNewEnv(env XLEnv) XLEnv {
	return XLEnv(NewXLPair(NewXLMap(),(*XLPair)(env)))
}

func PopEnv(env XLEnv) XLEnv {
	return XLEnv(env.Snd.(*XLPair))
}

func UnboxEnv(env XLEnv) *XLPair {
	return env
}
func RegisterNativeFunction(fname string, f XLNativeFunctionType, lazy bool, env XLEnv) {
	EnvPut(env, fname, NewXLNativeFunction(f, lazy))
}

func NativeApp(f XLNativeFunction, params XLObj, env XLEnv) (XLObj, bool) {
	return f.Func(params, env)
}

func Eval(params XLObj, env XLEnv) (XLObj, bool) {
	exp := params.(*XLPair).Fst
	//fmt.Println("Eval", exp)
	return ExpEval(exp, env)
}


func ParamsEval(params XLObj, env XLEnv) (XLObj, bool) {
	if params.XLObjType() == DT_Nil {
		return Nil, true
	}
	fst, _ := ExpEval(params.(*XLPair).Fst, env)
	rest, _ := ParamsEval(params.(*XLPair).Snd, env)
	return NewXLPair(fst, rest), true
}

func ParamsLazy(params XLObj, env XLEnv) (XLObj, bool) {
	if params.XLObjType() == DT_Nil {
		return Nil, true
	}
	fst := NewXLLazy(params.(*XLPair).Fst, env)
	rest, _ := ParamsLazy(params, env)
	return NewXLPair(fst, rest), true
}

func ExpEval(exp XLObj, env XLEnv) (XLObj, bool) {
	//fmt.Println("ExpEval: ", PrettyPrint(exp))
	switch exp.XLObjType() {
	case DT_Symbol:
		s := exp.(*XLSymbol).Value
		return EnvGet(env, s)
	case DT_Pair:
		f, _ := ExpEval(PairCar(exp), env)
		if f.XLObjType() == DT_NativeFunction {
			//fmt.Println("Native", PairCar(exp), PrettyPrint(PairCdr(exp)))
			var params XLObj
			if !f.(*XLNativeFunction).Lazy {
				params, _ = ParamsEval(PairCdr(exp), env)
				//fmt.Println("params", params, ok)
			} else {
				params = NewXLLazy(PairCdr(exp), env)
			}
			//fmt.Println("Native", PairCar(exp), PrettyPrint(params))

			return f.(*XLNativeFunction).Func(params, env)
		} else if f.XLObjType() == DT_Function {
			var params XLObj
			if !f.(*XLFunction).Lazy {
				params, _ = ParamsEval(PairCdr(exp), env)
			} else {
				params, _ = ParamsLazy(PairCdr(exp), env)
			}
			nenv := PushNewEnv(f.(*XLFunction).Env)
			pp := params
			//fmt.Println(f.(XLFunction).Params)
			//fmt.Println(PairCdr(exp))
			for _, s := range f.(*XLFunction).Params {
				EnvPut(nenv, s, PairCar(pp))
				pp = PairCdr(pp)
			}
			//fmt.Println(XLEnv(nenv).Fst)
			//fmt.Println(PrettyPrint(f.(XLFunction).Body))
			//fmt.Println(ExpEval(f.(XLFunction).Body, nenv))
			return ExpEval(f.(*XLFunction).Body, nenv)
		} else {
			return Nil, false
		}

	case DT_Lazy:
		l := exp.(*XLLazy)
		//fmt.Println(&l, l)
		if !l.Done {
			l.Value, _ = ExpEval(l.Value, l.Env)
			l.Done = true
		}
		return l.Value, true
	default:
		return exp, true
	}
}

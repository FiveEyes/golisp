package core

import "fmt"

func init() {	
	RegisterNativeFunction("lam", Lambda, true, BasicEnv)
	fmt.Println("lamNF.init done")
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

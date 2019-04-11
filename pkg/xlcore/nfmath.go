package core

import (
	"fmt"
)

func init() {
	RegisterNativeFunction("+", Add, false, BasicEnv)
	RegisterNativeFunction("-", Minus, false, BasicEnv)
	RegisterNativeFunction("eq?", Equal, false, BasicEnv)
	fmt.Println("mathNF.init done")
}

func addInt(obj XLObj, env XLEnv) (XLObj, bool) {
	sum := 0
	for p := obj; p.XLObjType() != DT_Nil; p = PairCdr(p) {
		sum += PairCar(p).(*XLInt).Value
	}
	return NewXLInt(sum), true
}

func addFloat(obj XLObj, env XLEnv) (XLObj, bool) {
	var sum float64 = 0
	for p := obj; p.XLObjType() != DT_Nil; p = PairCdr(p) {
		x := PairCar(p)
		if x.XLObjType() == DT_Int {
			sum += float64(PairCar(p).(*XLInt).Value)
		} else {
			sum += PairCar(p).(*XLFloat).Value
		}
	}
	return NewXLFloat(sum), true
}

func Add(obj XLObj, env XLEnv) (XLObj, bool) {
	flag := false
	for p := obj; p.XLObjType() != DT_Nil; p = PairCdr(p) {
		x := PairCar(p)
		if x.XLObjType() == DT_Int {
		} else if x.XLObjType() == DT_Float {
			flag = true
		} else {
			return Nil, false
		}
	}
	if !flag {
		return addInt(obj, env)
	} else {
		return addFloat(obj, env)
	}
}

func Minus(obj XLObj, env XLEnv) (XLObj, bool) {
	sum := PairCar(obj).(*XLInt).Value
	for p := PairCdr(obj); p.XLObjType() != DT_Nil; p = PairCdr(p) {
		sum -= PairCar(p).(*XLInt).Value
	}
	return NewXLInt(sum), true
}

func Equal(obj XLObj, env XLEnv) (XLObj, bool) {
	fst := PairCar(obj)
	snd := PairCar(PairCdr(obj))
	//fmt.Println(fst, snd, fst == snd)
	if fst.XLObjType() != snd.XLObjType() {
		return Zero, true
	}
	switch fst.XLObjType() {
	case DT_Int:
		if fst.(*XLInt).Value == snd.(*XLInt).Value {
			return One, true
		} else {
			return Zero, true
		}
	case DT_Float:
		if fst.(*XLFloat).Value == snd.(*XLFloat).Value {
			return One, true
		} else {
			return Zero, true
		}
	
	}
	if fst == snd {
		return NewXLInt(1), true
	} else {
		return NewXLInt(0), true
	}
}


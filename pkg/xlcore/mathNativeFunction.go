package core

import (
	_"fmt"
)

func init() {
	//RegisterNativeFunction("+", Add, false, BasicEnv)
	//RegisterNativeFunction("eq?", Equal, false, BasicEnv)
}

func addInt(obj XLObj, env XLEnv) (XLObj, bool) {
	sum := 0
	for p := obj; p.XLObjType() != DT_Nil; p = PairCdr(p) {
		sum += PairCar(p).(XLInt).Value
	}
	return XLInt {Value: sum}, true
}

func addFloat(obj XLObj, env XLEnv) (XLObj, bool) {
	var sum float64 = 0
	for p := obj; p.XLObjType() != DT_Nil; p = PairCdr(p) {
		x := PairCar(p)
		if x.XLObjType() == DT_Int {
			sum += float64(PairCar(p).(XLInt).Value)
		} else {
			sum += PairCar(p).(XLFloat).Value
		}
	}
	return XLFloat {Value: sum}, true
}

func Add(obj XLObj, env XLEnv) (XLObj, bool) {
	flag := false
	for p := obj; p.XLObjType() != DT_Nil; p = PairCdr(p) {
		x := PairCar(p)
		if x.XLObjType() == DT_Int {
		} else if x.XLObjType() == DT_Float {
			flag = true
		} else {
			return XLNil{}, false
		}
	}
	if !flag {
		return addInt(obj, env)
	} else {
		return addFloat(obj, env)
	}
}

func Minus(obj XLObj, env XLEnv) (XLObj, bool) {
	sum := PairCar(obj).(XLInt).Value
	for p := PairCdr(obj); p.XLObjType() != DT_Nil; p = PairCdr(p) {
		sum -= PairCar(p).(XLInt).Value
	}
	return XLInt {Value: sum}, true
}

func Equal(obj XLObj, env XLEnv) (XLObj, bool) {
	fst := PairCar(obj)
	snd := PairCar(PairCdr(obj))
	//fmt.Println(fst, snd, fst == snd)
	if fst == snd {
		return XLInt{Value:1}, true
	} else {
		return XLInt{Value:0}, true
	}
}

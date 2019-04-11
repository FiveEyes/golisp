package core

//XLEnv List XLMap
type XLEnv *XLPair

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
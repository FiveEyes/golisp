package core

func PairCar(p XLObj) XLObj {
	return p.(XLPair).Fst
}

func PairCdr(p XLObj) XLObj {
	return p.(XLPair).Snd
}

// list
func ListLength(p XLObj) int {
	l := 0
	for ; p.XLObjType() != DT_Nil; p = p.(XLPair).Snd {
		l += 1
	}
	return l
}
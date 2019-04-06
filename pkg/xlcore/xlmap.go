package core


func MapGet(m *XLMap, s string) (XLObj, bool) {
	if v, ok := m.Map[s]; ok {
		return v, true
	} else {
		return Nil, false
	}
}

func MapPut(m *XLMap, s string, v XLObj) {
	m.Map[s] = v
}

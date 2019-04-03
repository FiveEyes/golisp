package core

func NewMap() XLMap {
	return XLMap{Map:make(map[string]XLObj)}
}

func MapGet(m XLMap, s string) (XLObj, bool) {
	if v, ok := m.Map[s]; ok {
		return v, true
	} else {
		return nil, false
	}
}

func MapPut(m XLMap, s string, v XLObj) {
	m.Map[s] = v
}
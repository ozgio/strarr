package strarr

type none struct{}

func Merge(a []string, b []string) []string {
	if len(a) == 0 {
		return append([]string{}, b...)
	} else if len(b) == 0 {
		return append([]string{}, a...)
	}

	m := make(map[string]none)
	res := make([]string, 0, len(a)+len(b))
	res = append([]string{}, a...)

	for _, v := range a {
		m[v] = none{}
	}
	for _, v := range b {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

func Unique(a []string) []string {
	if len(a) == 0 {
		return []string{}
	}

	m := make(map[string]none)
	res := make([]string, 0, len(a))

	for _, v := range a {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = none{}
		}
	}

	return res
}

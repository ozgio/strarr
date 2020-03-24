package strarr

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

package strarr

func Exclude(target []string, arr []string) []string {
	if len(target) == 0 || len(arr) == 0 {
		return append([]string{}, target...)
	}

	m := make(map[string]bool)
	res := make([]string, 0, len(target))
	for _, v := range arr {
		m[v] = true
	}
	for _, v := range target {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

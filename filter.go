package strarr

func Filter(arr []string, fn func(string) bool) []string {
	res := make([]string, 0, len(arr))
	for _, v := range arr {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

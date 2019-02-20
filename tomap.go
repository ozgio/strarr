package strarr

func ToMap(arr []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	return m
}

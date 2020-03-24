package strarr

func Map(arr []string, fn func(string) string) {
	for i, v := range arr {
		arr[i] = fn(v)
	}
}

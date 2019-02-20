package strarr

func Map(arr []string, fn func(string) string) {
	for i, v := range arr {
		arr[i] = fn(v)
	}
}

func Filter(arr []string, fn func(string) bool) []string {
	res := make([]string, 0, len(arr))
	for _, v := range arr {
		if fn(v) {
			res = append(res, v)
		}
	}

	return res
}

// Any

// Every

// fill (pattern, count)

// Array.prototype.reverse()
// Array.prototype.shift()
// Array.prototype.slice()
// Array.prototype.some()
// Array.prototype.sort()
// Array.prototype.splice()

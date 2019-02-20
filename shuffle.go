package strarr

import "math/rand"

func Shuffle(arr []string) {
	for i := 0; i < len(arr); i++ {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

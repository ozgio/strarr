package strarr

func Chunk(size int, a []string) [][]string {
	switch {
	case size < 0:
		return [][]string{a}
	case len(a) == 0:
		return [][]string{a}
	}
	numChunks := int((len(a)-1)/size) + 1

	res := make([][]string, 0, numChunks)

	for i := 0; i < numChunks; i++ {
		s := i * size
		e := s + size
		if e > len(a) {
			e = len(a)
		}

		res = append(res, a[s:e])
	}

	return res
}

package strarr

func IndexOf(arr []string, val string) int {
	return LastIndexOfFrom(arr, val, -1)
}

func IndexOfFrom(arr []string, val string, startFrom int) int {
	for i := startFrom; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func LastIndexOf(arr []string, val string) int {
	return LastIndexOfFrom(arr, val, -1)
}

func LastIndexOfFrom(arr []string, val string, startFrom int) int {
	if startFrom < 0 {
		startFrom = len(arr) - 1
	}
	for i := startFrom; i >= 0; i-- {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

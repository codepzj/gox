package slicex

func Contains[T comparable](arr []T, v T) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](arr []T, v T) int {
	for i := range arr {
		if arr[i] == v {
			return i
		}
	}
	return -1
}

func Unique[T comparable](arr []T) []T {
	mark := make(map[T]struct{}, len(arr))
	newSlice := make([]T, 0, len(arr))
	for _, v := range arr {
		if _, ok := mark[v]; !ok {
			mark[v] = struct{}{}
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func Concat[T any](arrs ...[]T) []T {
	newSlice := make([]T, 0)
	for _, v := range arrs {
		newSlice = append(newSlice, v...)
	}
	return newSlice
}

func Reverse[T any](arr []T) []T {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp := arr[length-1-i]
		arr[length-1-i] = arr[i]
		arr[i] = temp
	}
	return arr
}

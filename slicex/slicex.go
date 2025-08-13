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

func Map[T any, R any](arr []T, fn func(T) R) []R {
	newSlice := make([]R, len(arr))
	for i, v := range arr {
		newSlice[i] = fn(v)
	}
	return newSlice
}

func Filter[T any](arr []T, fn func(T) bool) []T {
	newSlice := make([]T, 0)
	for _, v := range arr {
		if fn(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

// 把一系列元素“折叠”成一个值, 每次执行的结果与前一次得到的结果进行运算
func Reduce[T any, R any](arr []T, fn func(R, T) R, init R) R {
	for _, v := range arr {
		init = fn(init, v)
	}
	return init
}

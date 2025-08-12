package numberx

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

func Abs[T Number](a T) T {
	if a > 0 {
		return a
	}
	return -a
}

func Sum[T Number](nums ...T) T {
	var sum T = 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

func Average[T Number](nums ...T) float64 {
	return float64(Sum(nums...)) / float64(len(nums))
}

func IsBetween[T Number](x, min, max T) bool {
	return x >= min && x <= max
}

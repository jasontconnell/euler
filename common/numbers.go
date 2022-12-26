package common

import "math"

type Number interface {
	int | int64 | int32 | int16 | uint64 | uint32 | uint16
}

func IsPrime[N Number](n N) bool {
	if n == 1 || n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}
	for i := N(2); i < Sqrt(n)+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Sqrt[N Number](n N) N {
	return N(math.Sqrt(float64(n)))
}

func Max[N Number](cur N, check ...N) N {
	nmax := cur
	for _, n := range check {
		if n > nmax {
			nmax = n
		}
	}
	return nmax
}

func GCD[N Number](a, b N) N {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM[N Number](a, b N, nums ...N) N {
	result := a * b / GCD(a, b)

	for i := 0; i < len(nums); i++ {
		result = LCM(result, nums[i])
	}

	return result
}

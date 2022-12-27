package common

import (
	"math"
	"sort"
)

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

func Sum[N Number](list []N) N {
	var init N
	for _, n := range list {
		init += n
	}
	return init
}

func ProperDivisors[N Number](n N) []N {
	d := Divisors(n)
	return d[:len(d)-1]
}

func Divisors[N Number](n N) []N {
	if IsPrime(n) {
		return []N{1, n}
	}

	list := []N{}
	for i := N(1); i < Sqrt(n)+1; i++ {
		if n%i == 0 {
			list = append(list, i)
			if i != n/i {
				list = append(list, n/i)
			}
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	return list
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

func BigNumberCalc[N Number](initial N, max int, op func(idx int, n N) N) []N {
	digits := []N{initial}
	for i := 1; i < max; i++ {
		var carry N
		for j := len(digits) - 1; j >= 0; j-- {
			x := op(i, digits[j]) + carry
			digit := x % 10
			carry = x / 10
			digits[j] = digit
		}
		c := carry
		for c > 0 {
			digit := c % 10
			digits = append([]N{digit}, digits...)
			c /= 10
		}
	}
	return digits
}

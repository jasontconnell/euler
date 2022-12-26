package common

func IsPrime(n int) bool {
	if n == 1 {
		return true
	}
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

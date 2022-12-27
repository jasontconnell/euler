package common

func Permutate[T any](items []T) [][]T {
	var ret [][]T

	if len(items) == 2 {
		ret = append(ret, []T{items[0], items[1]})
		ret = append(ret, []T{items[1], items[0]})
	} else {

		for i := 0; i < len(items); i++ {
			cp := make([]T, len(items))
			copy(cp, items)

			t := cp[i]
			sh := append(cp[:i], cp[i+1:]...)
			perm := Permutate(sh)

			for _, p := range perm {
				p = append([]T{t}, p...)
				ret = append(ret, p)
			}
		}
	}

	return ret
}

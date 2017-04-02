package passgo

import "os"

func any(char byte, set []byte) bool {
	for _, c := range set {
		if char == c {
			return true
		}
	}
	return false
}

func all(word string, set []byte) bool {
	for _, char := range []byte(word) {
		if !any(char, set) {
			return false
		}
	}
	return true
}

func lcg(a, c, m, seed uint32) func() uint32 {
	r := seed
	return func() uint32 {
		r = (a*r + c) % m
		return r
	}
}

func pow10(pow int) int {
	res, val := 1, 10
	for pow > 0 {
		if pow&1 != 0 {
			res *= val
		}
		val *= val
		pow >>= 1 // equal `pow /= 2`
	}
	return res
}

func atoi(digit string) (r int) {
	p := pow10(len(digit) - 1)
	for _, c := range []byte(digit) {
		r += atoiMap[c] * p
		p /= 10
	}
	return
}

func randint(foo func() uint32, max int) func() int {
	return func() int {
		return int(foo()) % max
	}
}

func switchFilter(c byte) bool {
	b, ok := m[c]
	filter[b] = true
	return ok
}

func flagParse(flags []string) error {
	for _, a := range flags {
		charr := []byte(a)
		if charr[0] == '-' {
			if all(string(charr[0:2]), []byte("-")) {
				clean = clean && switchFilter(charr[2])
			} else {
				for _, c := range charr[1:] {
					clean = clean && switchFilter(c)
				}
			}
		} else if all(a, getter[2]) {
			l = atoi(a)
		}
	}
	return nil
}

// PassGen ... READ FLAGS AND GENERATING PASS
func PassGen(flags []string) string {
	alp := &alpha
	flagParse(flags)
	for i, b := range filter {
		if b {
			*alp = append(alpha, getter[i]...)
		}
	}
	rnd := randint(lcg(1103515245, 12345, uint32((1<<31)-1), uint32(os.Getppid()+os.Getpid())), len(alpha))
	for ; l > 0; l-- {
		pass += string(alpha[rnd()])
	}
	return pass
}

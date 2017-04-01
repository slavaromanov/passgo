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

func pow10(b int) int {
	p, a := 1, 10
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
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

func defaultAlpha() []byte {
	return append(getter[0], append(getter[1], getter[2]...)...)
}

func compare(x ...string) (one string) {
	for _, s := range x {
		one += s
	}
	return
}

func switchFilter(c byte) string {
	b, ok := m[c]
	if ok {
		filter[b] = true
		return ""
	}
	return "Invalid option\n"
}

func flagParse(flags []string) (res string) {
	for _, a := range flags {
		charr := []byte(a)
		if charr[0] == '-' {
			if all(string(charr[0:2]), []byte("-")) {
				res += switchFilter(charr[2])
			} else {
				for _, c := range charr[1:] {
					res += switchFilter(c)
				}
			}
		} else if all(a, getter[2]) {
			l = atoi(a)
		}
	}
	return
}

// PassGen ... READ FLAGS AND GENERATING PASS
func PassGen(flags []string) string {
	alp := &alpha
	fr := flagParse(flags)
	if fr != "" {
		return fr
	}
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

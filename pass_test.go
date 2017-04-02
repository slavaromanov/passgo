package passgo

import "testing"

func TestAny(t *testing.T) {
	if any('c', []byte("klasjlkcaw")) == false {
		t.Fatal("TestAny failed")
	}
}

func TestAll(t *testing.T) {
	if all("hello", []byte("abcdefghijklmnop")) == false {
		t.Fatal("TestAll failed")
	}
}

func TestPow10(t *testing.T) {
	if pow10(5) != 10*10*10*10*10 {
		t.Fail()
	}
}

func TestAtoi(t *testing.T) {
	if atoi("123456") != 123456 {
		t.Fail()
	}
}

func TestRandint(t *testing.T) {
	r := randint(lcg(1103515245, 12345, uint32((1<<31)-1), uint32(0)), 16)
	rr := r()
	if rr > 15 || rr < 0 {
		t.Fail()
	}
}

func TestPassGen(t *testing.T) {
	PassGen([]string{"10", "-d"})
	if pass == "" {
		t.Fail()
	}
}

func TestFlagParse(t *testing.T) {
	if flagParse([]string{"10", "--dige", "--spec", "-lu"}) != nil && clean {
		t.Fail()
	}
}

func TestSwitchFilter(t *testing.T) {
	if !switchFilter('d') || filter[m['d']] != true {
		t.Fail()
	}
}

func TestM(t *testing.T) {
	if !t.Run("string", func(t *testing.T) {
		main()
	}) {
		t.Fail()
	}
}

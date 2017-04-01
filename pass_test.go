package passgo

import "testing"

func TestAny(t *testing.T) {
	if any('c', []byte("klasjlkcaw")) == false {
		t.Fatal("TestAny failed")
	}
	t.Log("TestAny pass")
}

func TestAll(t *testing.T) {
	if all("hello", []byte("abcdefghijklmnop")) == false {
		t.Fatal("TestAll failed")
	}
	t.Log("TestAny pass")
}

func TestPow10(t *testing.T) {
	if pow10(5) != 100000 {
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

func TestDefaultAlpha(t *testing.T) {
	if !all(string(defaultAlpha()), alphabeth) {
		t.Fail()
	}
}

func TestFlagParse(t *testing.T) {
	fp := flagParse([]string{"10", "--dige", "--spec", "-lu"})
	t.Log(fp)
	if fp != "" {
		t.Fail()
	}
}

func TestCompare(t *testing.T) {
	if compare("test ", "pass!") != "test pass!" {
		t.Fail()
	}
}

func TestSwitchFilter(t *testing.T) {
	lg := switchFilter('d')
	if lg != "" || filter[m['d']] != true {
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

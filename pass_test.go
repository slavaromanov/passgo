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
	r := randint(lcg(), 16)
	rr := r()
	if rr > 15 || rr < 0 {
		t.Fail()
	}
}

func TestGenAlpha(t *testing.T) {
	if len(GetAlphabet([]string{"10", "--dige", "--spec", "-lu"})) <= 0 {
		t.Fail()
	}
}

func TestSwitchFilter(t *testing.T) {
	if !switchFilter('l') {
		t.Fail()
	}
}

func TestSerializeMap(t *testing.T) {
	if len(serializeMap()[false]) != 62 {
		t.Fail()
	}
}

func TestSerializer(t *testing.T) {
	if string(serializer(map[byte]bool{
		'f': false,
	})) != "f" {
		t.Fail()
	}
}

func TestNewPass(t *testing.T) {
	if len(NewPass([]string{"16"})) != 16 {
		t.Fail()
	}
}

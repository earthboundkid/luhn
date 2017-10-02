package luhn

import (
	"testing"
)

var testcases = []struct {
	string
	int
	valid  bool
	checkb byte
	checki int
}{
	{"5", 5, false, '9', 9},
	{"059", 59, true, '6', 6},
	{"7992739871", 7992739871, false, '3', 3},
	{"79927398713", 79927398713, true, '8', 8},
	{"4992739871", 4992739871, false, '6', 6},
	{"49927398716", 49927398716, true, '8', 8},
	{"468192572269901", 468192572269901, false, '0', 0},
	{"4681925722699010", 4681925722699010, true, '9', 9},
	{"123456781234567", 123456781234567, false, '0', 0},
}

func TestCheckDigit(t *testing.T) {
	for _, test := range testcases {
		test := test
		var r byte
		t.Run(test.string, func(t *testing.T) {
			r = CheckDigit(test.string)
			if r != test.checkb {
				t.Errorf("Failure for %s: expected %d; got %s",
					test.string, test.checki, string(r))
			}
		})
	}
}

func TestCheckDigitInt(t *testing.T) {
	for _, test := range testcases {
		test := test
		var r int
		t.Run(test.string, func(t *testing.T) {
			r = CheckDigitInt(test.int)
			if r != test.checki {
				t.Errorf("Failure for %s: expected %d; got %d",
					test.string, test.checki, r)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	for _, test := range testcases {
		test := test
		var r bool
		t.Run(test.string, func(t *testing.T) {
			r = IsValid(test.string)
			if r != test.valid {
				t.Errorf("Failure for %s: expected valid == %t",
					test.string, test.valid)
			}
		})
	}
}

func TestIsValidInt(t *testing.T) {
	for _, test := range testcases {
		test := test
		var r bool
		t.Run(test.string, func(t *testing.T) {
			r = IsValidInt(test.int)
			if r != test.valid {
				t.Errorf("Failure for %s: expected valid == %t",
					test.string, test.valid)
			}
		})
	}
}

func BenchmarkCheckDigit(b *testing.B) {
	for _, test := range testcases {
		test := test
		var r byte
		b.Run(test.string, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = CheckDigit(test.string)
			}
		})
	}
}

func BenchmarkCheckDigitInt(b *testing.B) {
	for _, test := range testcases {
		test := test
		var r int
		b.Run(test.string, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = CheckDigitInt(test.int)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	for _, test := range testcases {
		test := test
		var r bool
		b.Run(test.string, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = IsValid(test.string)
			}
		})
	}
}

func BenchmarkIsValidInt(b *testing.B) {
	for _, test := range testcases {
		test := test
		var r bool
		b.Run(test.string, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r = IsValidInt(test.int)
			}
		})
	}
}

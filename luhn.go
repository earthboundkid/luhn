// Package luhn implements the Luhn algorithm, which is commonly used to add
// checksums to credit cards and identity codes.
//
// See https://en.wikipedia.org/wiki/Luhn_algorithm
package luhn

func sum(n int, parity bool) (total int) {
	for n > 0 {
		d := n % 10
		if parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		total += d
		parity = !parity
		n /= 10
	}
	return total
}

// IsValidInt returns true if the input is valid according to the Luhn algorithm.
func IsValidInt(n int) bool {
	return sum(n, false)%10 == 0
}

// CheckDigitInt returns the digit needed to make its input valid.
func CheckDigitInt(n int) int {
	return sum(n, true) * 9 % 10
}

func sumb(b []byte, parity bool) (total int) {
	for n := len(b) - 1; n >= 0; n-- {
		d := b[n] - '0'
		if parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}

		total += int(d)
		parity = !parity
	}
	return total
}

// IsValid returns true if the input is valid according to the Luhn algorithm.
//
// Note that IsValid assumes input characters are all digits. Output is
// undefined for unexpected characters.
func IsValid(s string) bool {
	return sumb([]byte(s), false)%10 == 0
}

// CheckDigit returns the digit needed to make its input valid.
func CheckDigit(s string) byte {
	return byte(sumb([]byte(s), true)*9%10) + '0'
}

package pkg

import (
	"fmt"

	"github.com/bytedance/gopkg/lang/fastrand"
)

// verifyingDigits generates the two verifying digits (p3 and p4) for CPF,
// as stated by Receita Federal of Brazil.
func verifyingDigits(p1, p2 uint32) (p3, p4 uint32) {
	if p1 >= 100000000 {
		return 0, 0
	}

	var tmp, i uint32

	p3 = p2 * 9
	p4 = p2 * 8

	tmp = p1
	i = 1
	for i < 10 {
		d := tmp % 10
		p3 += d * (9 - i)
		p4 += d * (9 - ((i + 1) % 10))
		tmp = tmp / 10
		i++
	}

	p3 = (p3 % 11) % 10
	p4 += p3 * 9
	p4 = (p4 % 11) % 10

	return p3, p4
}

// CPF generates n strings with valid CPF numbers.
func CPF(n int) []string {
	r := make([]string, n)

	for i := range n {
		p1 := fastrand.Uint32n(100000000) // Part 1: random 8 digit number
		p2 := fastrand.Uint32n(10)        // Part 2: tax region code
		p3, p4 := verifyingDigits(p1, p2) // Part 3: verifying digits
		r[i] = fmt.Sprintf("%d%d%d%d", p1, p2, p3, p4)
	}

	return r
}
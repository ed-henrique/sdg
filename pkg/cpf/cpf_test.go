package pkg

import (
	"fmt"
	"testing"
)

func TestVerifyingDigits(t *testing.T) {
	t.Parallel()

	t.Run("cpf within range [0, 100000000)", func(t *testing.T) {
		d1, d2 := verifyingDigits(1234567, 8)
		var r1, r2 uint32 = 9, 0

		if d1 != r1 || d2 != r2 {
			t.Errorf("got %d%d r %d%d", d1, d2, r1, r2)
		}
	})

	t.Run("cpf out of bounds", func(t *testing.T) {
		d1, d2 := verifyingDigits(101234567, 8)
		var r1, r2 uint32 = 0, 0

		if d1 != r1 || d2 != r2 {
			t.Errorf("got %d%d r %d%d", d1, d2, r1, r2)
		}
	})
}

var table = []int{
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
}

func BenchmarkCPF(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CPF(v)
			}
		})
	}
}

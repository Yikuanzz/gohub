package benchmark

import (
	"fmt"
	"testing"
)

var numbers = []int{
	99,
	999,
	99999,
	999999,
}

func BenchmarkPrimeNumbers(b *testing.B) {
	for _, n := range numbers {
		b.Run(fmt.Sprintf("calc_prime_number_from 2 to %d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				primeNumbers(n)
			}
		})
	}
}

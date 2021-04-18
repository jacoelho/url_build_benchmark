package url_performance

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprinfString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s%s%s", data[0], data[1], data[2])
		_ = s
	}
}

func BenchmarkSprinfDigit(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s%s%d", data[0], data[1], i)
		_ = s
	}
}

func BenchmarkSprinfDigitItoa(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s%s%s", data[0], data[1], strconv.Itoa(digit))
		_ = s
	}
}

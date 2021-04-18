package url_performance

import (
	"strconv"
	"testing"
)

func BenchmarkConcatString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := data[0] + data[1] + data[2]
		_ = s
	}
}

func BenchmarkConcatStringAndItoa(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		s := data[0] + data[1] + strconv.Itoa(digit)
		_ = s
	}
}

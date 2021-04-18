package url_performance

import (
	"bytes"
	"strconv"
	"testing"
)

func BenchmarkBytesString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b := new(bytes.Buffer)

		_, _ = b.WriteString(data[0])
		_, _ = b.WriteString(data[1])
		_, _ = b.WriteString(data[2])
		s := b.String()
		_ = s
	}
}

func BenchmarkBytesStringAndItoa(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b := new(bytes.Buffer)

		_, _ = b.WriteString(data[0])
		_, _ = b.WriteString(data[1])
		_, _ = b.WriteString(strconv.Itoa(digit))
		s := b.String()
		_ = s
	}
}

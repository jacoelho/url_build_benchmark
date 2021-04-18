package url_performance

import (
	"net/url"
	"testing"
)

func BenchmarkURLParseString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		u, _ := url.Parse(data[0])
		p1, _ := url.Parse(data[1])
		p2, _ := url.Parse(data[2])

		s := u.ResolveReference(p1).ResolveReference(p2).String()
		_ = s
	}
}

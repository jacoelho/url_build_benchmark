package url_performance

import (
	"net/url"
	"testing"

	"github.com/google/go-querystring/query"
)

func BenchmarkURLParseString(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		u, _ := url.Parse(urlString)
		s := u.String()
		_ = s
	}
}

func BenchmarkURLParseResolveReference(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		u, _ := url.Parse(data[0])
		p1, _ := url.Parse(data[1])
		p2, _ := url.Parse(data[2])

		s := u.ResolveReference(p1).ResolveReference(p2).String()
		_ = s
	}
}

func BenchmarkURLQueryEncode(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		u, _ := url.Parse(urlString)

		query := url.Values{}
		query.Set("filter[page]", "1")
		u.RawQuery = query.Encode()
		s := u.String()
		_ = s
	}
}

func BenchmarkURLQueryEncodePackageQueryString(b *testing.B) {
	b.ReportAllocs()

	type Options struct {
		Page string `url:"filter[page]"`
	}

	for i := 0; i < b.N; i++ {
		u, _ := url.Parse(urlString)

		q, _ := query.Values(Options{Page: "1"})
		u.RawQuery = q.Encode()
		s := u.String()
		_ = s
	}
}

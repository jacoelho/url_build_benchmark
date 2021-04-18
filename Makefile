bench:
	go test -bench=.

bench_10s:
	go test -bench=. -benchtime=10s

bench-escape:
	go test -bench=.  -gcflags="-m"

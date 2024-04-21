# Servers benchmark

http servers performance test

- http standard library
- gorilla mux
- echo
- gin

## scripts

> run benchmarks
> ```bash
> $ cd api
> ```
> ```bash
> $ go test -bench . -benchtime=1000x -benchmem
> ```
> ```
> ...
> goos: linux
> goarch: amd64
> pkg: github.com/edwintrumpet/servers-benchmark/api
> cpu: 12th Gen Intel(R) Core(TM) i5-12500H
> BenchmarkMuxParallel-16             1000            181231 ns/op           29062 B/op        150 allocs/op
> BenchmarkStandarParallel-16         1000             67661 ns/op           27949 B/op        143 allocs/op
> BenchmarkGinParallel-16             1000             69080 ns/op           28337 B/op        147 allocs/op
> BenchmarkEchoParallel-16            1000             14002 ns/op           28488 B/op        148 allocs/op
> PASS
> ok      github.com/edwintrumpet/servers-benchmark/api   1.432s
> ```


> [!WARNING]
> the order of the benchmarks affects the result.  
> the first benchmark takes the most time and the last benchmark takes the least
> amount of time


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
> $ go test -bench .
> ```

> [!CAUTION]
> run benchmarks returns an error
> ```
> goos: linux
> goarch: amd64
> pkg: github.com/edwintrumpet/servers-benchmark/api
> cpu: 12th Gen Intel(R) Core(TM) i5-12500H
> BenchmarkStandarParallel-16        16416             85128 ns/op
> BenchmarkMuxParallel-16            86035           1700906 ns/op
> [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
>
> [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
>  - using env:   export GIN_MODE=release
>  - using code:  gin.SetMode(gin.ReleaseMode)
>
> [GIN-debug] GET    /ping                     --> github.com/edwintrumpet/servers-benchmark/api/gin.Start.func1 (3 handlers)
> [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
> Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
> [GIN-debug] Listening and serving HTTP on :3002
> [GIN] 2024/04/20 - 12:28:28 | 200 |      94.222µs |       127.0.0.1 | GET      "/ping"
> BenchmarkGinParallel-16         [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
>
> [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
>  - using env:   export GIN_MODE=release
>  - using code:  gin.SetMode(gin.ReleaseMode)
>
> [GIN-debug] GET    /ping                     --> github.com/edwintrumpet/servers-benchmark/api/gin.Start.func1 (3 handlers)
> [GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
> Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
> [GIN-debug] Listening and serving HTTP on :3002
> [GIN-debug] [ERROR] listen tcp :3002: bind: address already in use
> 2024/04/20 12:28:28 listen tcp :3002: bind: address already in use
> exit status 1
> FAIL    github.com/edwintrumpet/servers-benchmark/api   158.299s
> ```

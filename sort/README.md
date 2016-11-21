# Sort

The followings are benchmarking result.

|name|times|speed|allocs|allocs|
| --- | --- | --- | --- | --- |
|BenchmarkSort1K/Quick-4 |   30000|     48905 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Shell-4 |   20000|     87715 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Heap-4  |   10000|    100674 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Bucket-4         |   10000|    167493 ns/op|   36766 B/op|     781 allocs/op|
|BenchmarkSort1K/Std-4            |   10000|    116066 ns/op|      32 B/op|       1 allocs/op|
|BenchmarkSort1K/Insert-4         |    3000|    408920 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Select-4         |    2000|   1014113 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Bubble-4         |    2000|   1181917 ns/op|       0 B/op|       0 allocs/op|


## Tips

To compare performance use [`benchcmp`](https://godoc.org/golang.org/x/tools/cmd/benchcmp)

```bash
$ go get golang.org/x/tools/cmd/benchcmp
$ go test -bench . -benchmem -cpuprofile cpu.prof > old.txt
$ go test -bench . -bemchmem -cpuprofile cpu.prof > new.txt
$ benchcmp old.txt new.txt
```

To check cpu profile,

```bash
$ go test -bench . -benchmem -cpuprofile cpu.prof
$ go tool pprof sort.test cpu.prof
```

(Why `sort.test` is required ? To find simbol localtion)

To check memory allocation,

```bash
$ go test -bench . -benchmem -memprofile mem.prof
$ go tool pprof -alloc_objects sort.test mem.prof
```


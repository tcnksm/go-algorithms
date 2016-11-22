# Sort

The followings are benchmarking result.

| name | times | speed | allocs | allocs |
| ---- | ----- | ----- | ------ | ------ |
|BenchmarkSort1K/Quick-4         |   30000|     47242 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Shell-4         |   20000|     88897 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Merge-4         |   20000|     96350 ns/op|   81920 B/op|    1023 allocs/op|
|BenchmarkSort1K/Heap-4          |   10000|    100483 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Bucket-4        |   10000|    169949 ns/op|   36761 B/op|     781 allocs/op|
|BenchmarkSort1K/Std-4           |   10000|    111402 ns/op|      32 B/op|       1 allocs/op|
|BenchmarkSort1K/Insert-4        |    3000|    408089 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Select-4        |    2000|   1017915 ns/op|       0 B/op|       0 allocs/op|
|BenchmarkSort1K/Bubble-4        |    2000|   1163597 ns/op|       0 B/op|       0 allocs/op|

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

To generate markdown table,

```bash
$ go get github.com/tcnksm/misc/cmd/benchtable
$ go test -bench . -benchmem | benchtable
```

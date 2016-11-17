
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

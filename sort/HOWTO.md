
To compare performance use [`benchcmp`](https://godoc.org/golang.org/x/tools/cmd/benchcmp)

```bash
$ go get golang.org/x/tools/cmd/benchcmp
$ go test -bench . > old.txt
$ go test -bench . > new.txt
$ benchcmp old.txt new.txt
```

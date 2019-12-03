Test of https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction in go:
```
[ivan@Ivans-MacBook-Pro ~/Documents/workspace/go-lang/src/c1rno/fast-sharding]# go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/c1rno/fast-sharding
BenchmarkModulo10-8                     47326318                24.7 ns/op
BenchmarkModulo100-8                    43543879                27.2 ns/op
BenchmarkModulo1000-8                   45278859                28.1 ns/op
BenchmarkModulo10000-8                  42604928                24.6 ns/op
BenchmarkModulo2147483647-8             42489333                25.3 ns/op
BenchmarkPowerWithShift10-8             61684336                18.5 ns/op
BenchmarkPowerWithShift100-8            61681197                18.3 ns/op
BenchmarkPowerWithShift1000-8           61412740                18.3 ns/op
BenchmarkPowerWithShift10000-8          61774681                18.3 ns/op
BenchmarkPowerWithShift2147483647-8     57444783                18.4 ns/op
PASS
ok      github.com/c1rno/fast-sharding  12.731s
```

It's obvious that `Power` faster than `Modulo`, but I'm not undestand, why it shows so bad distribution:

```
BenchmarkModulo10-8
Total length 10
1) key=0, value=10
2) key=1, value=10
3) key=2, value=10
4) key=4, value=10
5) key=7, value=10

Total length 10
1) key=8, value=1000
2) key=0, value=1000
3) key=2, value=1000
4) key=3, value=1000
5) key=4, value=1000

Total length 10
1) key=4, value=100000
2) key=5, value=100000
3) key=6, value=100000
4) key=8, value=100000
5) key=0, value=100000

Total length 10
1) key=6, value=4162012
2) key=5, value=4162012
3) key=0, value=4162012
4) key=1, value=4162012
5) key=2, value=4162012
```

vs

```
BenchmarkPowerWithShift10-8
Total length 1
1) key=0, value=100

Total length 1
1) key=0, value=10000

Total length 1
1) key=0, value=1000000

Total length 1
1) key=0, value=58990282

Total length 1
1) key=0, value=1

BenchmarkPowerWithShift100-8
Total length 1
1) key=0, value=100

Total length 1
1) key=0, value=10000

Total length 1
1) key=0, value=1000000

Total length 1
1) key=0, value=55563534

Total length 1
1) key=0, value=1

BenchmarkPowerWithShift1000-8
Total length 1
1) key=0, value=100

Total length 1
1) key=0, value=10000

Total length 1
1) key=0, value=1000000

Total length 1
1) key=0, value=59547790

Total length 1
1) key=0, value=1

BenchmarkPowerWithShift10000-8
Total length 1
1) key=0, value=100

Total length 1
1) key=0, value=10000

Total length 1
1) key=0, value=1000000

Total length 1
1) key=0, value=58366558

Total length 1
1) key=0, value=1

BenchmarkPowerWithShift2147483647-8
Total length 1
1) key=0, value=100

Total length 1
1) key=0, value=10000

Total length 1
1) key=0, value=1000000

Total length 1
1) key=0, value=56819101
```

which is look like all values fall into same key.

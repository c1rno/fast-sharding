package sharding

import (
	"fmt"
	"sort"
	"testing"
)

const length = 10

func BenchmarkModulo10(b *testing.B)         { bench(Modulo, b) }
func BenchmarkModulo100(b *testing.B)        { bench(Modulo, b) }
func BenchmarkModulo1000(b *testing.B)       { bench(Modulo, b) }
func BenchmarkModulo10000(b *testing.B)      { bench(Modulo, b) }
func BenchmarkModulo2147483647(b *testing.B) { bench(Modulo, b) }

func BenchmarkPowerWithShift10(b *testing.B)         { bench(PowerWithShift, b) }
func BenchmarkPowerWithShift100(b *testing.B)        { bench(PowerWithShift, b) }
func BenchmarkPowerWithShift1000(b *testing.B)       { bench(PowerWithShift, b) }
func BenchmarkPowerWithShift10000(b *testing.B)      { bench(PowerWithShift, b) }
func BenchmarkPowerWithShift2147483647(b *testing.B) { bench(PowerWithShift, b) }

func bench(algorithm CalculationAlgorithm, b *testing.B) {
	sh := Sharder(algorithm, length)
	m := make(map[uint32]int, length)
	var (
		ret uint32
		ok  bool
	)
	for n := 0; n < b.N; n++ {
		ret = sh(uint32(n))
		if _, ok = m[ret]; ok {
			m[ret] += 1
		} else {
			m[ret] = 1
		}
	}
	b.StopTimer()
	print5MostCommonItems(m)
}

func print5MostCommonItems(m map[uint32]int) {
	ss := make([]kv, 0, len(m))
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Printf("\nTotal length %d\n", len(ss))
	for index, kv := range ss {
		if index > 4 {
			break
		}
		fmt.Printf("%d) key=%d, value=%d\n", index+1, kv.Key, kv.Value)
	}
}

type kv struct {
	Key   uint32
	Value int
}

// Benchmarking
// $ go test -bench (running)
// $ go test -bench -benchtime 10s (targeting specified time)

// Others
// go test -benchmem (memory allocation)
// go test -trace {trace.out} (record execution trade for analysis)
// go test -{type}profile {file} (generate profile of requested type)
//    types: block, cover, cpu, mem, mutex

package benchmarking

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Marry had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Marry had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Marry had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}

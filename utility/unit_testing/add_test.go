package unit_testing

import (
	"testing"
)

// go test ./utility/unit_testing
func TestAdd(t *testing.T) {
	result := Add(2, 3)

	if result != 5 {
		t.Errorf("expected 5 but got %d", result)
	}
}

// go test ./utility/unit_testing -bench=.
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}

// BenchmarkAdd-10  1000000000  0.2253 ns/op
// 1. BenchmarkAdd-10 ==> GOMAXPROCS / số logical CPU được dùng (10 logical threads)
// 2. 1000000000 ==> số lần function đc run (1B times)
// 3. 0.2253 ns/op ==> 0.2253 nanoseconds/operation (mỗi lần run)

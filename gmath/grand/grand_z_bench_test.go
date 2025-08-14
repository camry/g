package grand_test

import (
    "testing"
    "time"

    "github.com/camry/g/gmath/grand"
)

// go test -bench=BenchmarkGRand_RangeInt -benchmem -count=10
func BenchmarkGRand_RangeInt(b *testing.B) {
    r := grand.NewRand(time.Now().UnixNano())
    for i := 0; i < b.N; i++ {
        r.RangeInt(1, 100000)
    }
}

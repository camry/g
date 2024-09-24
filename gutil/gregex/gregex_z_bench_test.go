package gregex_test

import (
    "regexp"
    "testing"

    "github.com/camry/g/gutil/gregex"
)

var pattern = `(\w+).+\-\-\s*(.+)`

var src = `G is best! -- Camry`

func Benchmark_IsMatchString(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gregex.IsMatchString(pattern, src)
    }
}

func Benchmark_MatchString(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gregex.MatchString(pattern, src)
    }
}

func Benchmark_Compile(b *testing.B) {
    var wcdRegexp = regexp.MustCompile(pattern)
    for i := 0; i < b.N; i++ {
        wcdRegexp.MatchString(src)
    }
}

func Benchmark_Compile_Actual(b *testing.B) {
    for i := 0; i < b.N; i++ {
        wcdRegexp := regexp.MustCompile(pattern)
        wcdRegexp.MatchString(src)
    }
}

package gerror_test

import (
    "errors"
    "testing"

    "github.com/camry/g/v2/gerrors/gcode"
    "github.com/camry/g/v2/gerrors/gerror"
)

var (
    // Wrap* 函数基准测试的基本错误。
    baseError = errors.New("test")
)

func Benchmark_New(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.New("test")
    }
}

func Benchmark_Newf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.Newf("%s", "test")
    }
}

func Benchmark_Wrap(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.Wrap(baseError, "test")
    }
}

func Benchmark_Wrapf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.Wrapf(baseError, "%s", "test")
    }
}

func Benchmark_NewSkip(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.NewSkip(1, "test")
    }
}

func Benchmark_NewSkipf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.NewSkipf(1, "%s", "test")
    }
}

func Benchmark_NewCode(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.NewCode(gcode.New(500, "", nil), "test")
    }
}

func Benchmark_NewCodef(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.NewCodef(gcode.New(500, "", nil), "%s", "test")
    }
}

func Benchmark_NewCodeSkip(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.NewCodeSkip(gcode.New(1, "", nil), 500, "test")
    }
}

func Benchmark_NewCodeSkipf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.NewCodeSkipf(gcode.New(1, "", nil), 500, "%s", "test")
    }
}

func Benchmark_WrapCode(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.WrapCode(gcode.New(500, "", nil), baseError, "test")
    }
}

func Benchmark_WrapCodef(b *testing.B) {
    for i := 0; i < b.N; i++ {
        gerror.WrapCodef(gcode.New(500, "", nil), baseError, "test")
    }
}

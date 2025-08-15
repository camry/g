package mutex_test

import (
    "testing"

    "github.com/camry/g/v2/gos/mutex"
)

var (
    safeLock   = mutex.New(false)
    unsafeLock = mutex.New(true)
)

func Benchmark_Safe_LockUnlock(b *testing.B) {
    for i := 0; i < b.N; i++ {
        safeLock.Lock()
        safeLock.Unlock()
    }
}

func Benchmark_UnSafe_LockUnlock(b *testing.B) {
    for i := 0; i < b.N; i++ {
        unsafeLock.Lock()
        unsafeLock.Unlock()
    }
}

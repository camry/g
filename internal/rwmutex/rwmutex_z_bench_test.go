package rwmutex_test

import (
    "testing"

    "github.com/camry/g/internal/rwmutex"
)

var (
    safeLock   = rwmutex.New(true)
    unsafeLock = rwmutex.New(false)
)

func Benchmark_Safe_LockUnlock(b *testing.B) {
    for i := 0; i < b.N; i++ {
        safeLock.Lock()
        safeLock.Unlock()
    }
}

func Benchmark_Safe_RLockRUnlock(b *testing.B) {
    for i := 0; i < b.N; i++ {
        safeLock.RLock()
        safeLock.RUnlock()
    }
}

func Benchmark_UnSafe_LockUnlock(b *testing.B) {
    for i := 0; i < b.N; i++ {
        unsafeLock.Lock()
        unsafeLock.Unlock()
    }
}

func Benchmark_UnSafe_RLockRUnlock(b *testing.B) {
    for i := 0; i < b.N; i++ {
        unsafeLock.RLock()
        unsafeLock.RUnlock()
    }
}

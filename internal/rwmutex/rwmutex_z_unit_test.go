package rwmutex_test

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/internal/rwmutex"
)

func TestRWMutexIsSafe(t *testing.T) {
    lock := rwmutex.New()
    assert.Equal(t, lock.IsSafe(), false)

    lock = rwmutex.New(false)
    assert.Equal(t, lock.IsSafe(), false)

    lock = rwmutex.New(false, false)
    assert.Equal(t, lock.IsSafe(), false)

    lock = rwmutex.New(true, false)
    assert.Equal(t, lock.IsSafe(), true)

    lock = rwmutex.New(true, true)
    assert.Equal(t, lock.IsSafe(), true)

    lock = rwmutex.New(true)
    assert.Equal(t, lock.IsSafe(), true)
}

func TestSafeRWMutex(t *testing.T) {
    var (
        localSafeLock = rwmutex.New(true)
        array         = make([]any, 0, 10)
    )

    go func() {
        localSafeLock.Lock()
        array = append(array, 1)
        time.Sleep(1000 * time.Millisecond)
        array = append(array, 1)
        localSafeLock.Unlock()
    }()
    go func() {
        time.Sleep(100 * time.Millisecond)
        localSafeLock.Lock()
        array = append(array, 1)
        time.Sleep(2000 * time.Millisecond)
        array = append(array, 1)
        localSafeLock.Unlock()
    }()
    time.Sleep(500 * time.Millisecond)
    assert.Equal(t, len(array), 1)
    time.Sleep(800 * time.Millisecond)
    assert.Equal(t, len(array), 3)
    time.Sleep(1000 * time.Millisecond)
    assert.Equal(t, len(array), 3)
    time.Sleep(1000 * time.Millisecond)
    assert.Equal(t, len(array), 4)
}

func TestSafeReaderRWMutex(t *testing.T) {
    var (
        localSafeLock = rwmutex.New(true)
        array         = make([]any, 0, 10)
    )
    go func() {
        localSafeLock.RLock()
        array = append(array, 1)
        time.Sleep(1000 * time.Millisecond)
        array = append(array, 1)
        localSafeLock.RUnlock()
    }()
    go func() {
        time.Sleep(100 * time.Millisecond)
        localSafeLock.RLock()
        array = append(array, 1)
        time.Sleep(2000 * time.Millisecond)
        array = append(array, 1)
        time.Sleep(1000 * time.Millisecond)
        array = append(array, 1)
        localSafeLock.RUnlock()
    }()
    go func() {
        time.Sleep(500 * time.Millisecond)
        localSafeLock.Lock()
        array = append(array, 1)
        localSafeLock.Unlock()
    }()
    time.Sleep(500 * time.Millisecond)
    assert.Equal(t, len(array), 2)
    time.Sleep(1000 * time.Millisecond)
    assert.Equal(t, len(array), 3)
    time.Sleep(1000 * time.Millisecond)
    assert.Equal(t, len(array), 4)
    time.Sleep(1000 * time.Millisecond)
    assert.Equal(t, len(array), 6)
}

func TestUnsafeRWMutex(t *testing.T) {
    var (
        localUnsafeLock = rwmutex.New()
        array           = make([]any, 0, 10)
    )
    go func() {
        localUnsafeLock.Lock()
        array = append(array, 1)
        time.Sleep(2000 * time.Millisecond)
        array = append(array, 1)
        localUnsafeLock.Unlock()
    }()
    go func() {
        time.Sleep(500 * time.Millisecond)
        localUnsafeLock.Lock()
        array = append(array, 1)
        time.Sleep(500 * time.Millisecond)
        array = append(array, 1)
        localUnsafeLock.Unlock()
    }()
    time.Sleep(800 * time.Millisecond)
    assert.Equal(t, len(array), 2)
    time.Sleep(800 * time.Millisecond)
    assert.Equal(t, len(array), 3)
    time.Sleep(200 * time.Millisecond)
    assert.Equal(t, len(array), 3)
    time.Sleep(500 * time.Millisecond)
    assert.Equal(t, len(array), 4)
}

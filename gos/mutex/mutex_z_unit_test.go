package mutex_test

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gos/mutex"
)

func TestMutexIsSafe(t *testing.T) {
    lock := mutex.New()
    assert.Equal(t, lock.IsSafe(), false)

    lock = mutex.New(false)
    assert.Equal(t, lock.IsSafe(), false)

    lock = mutex.New(false, false)
    assert.Equal(t, lock.IsSafe(), false)

    lock = mutex.New(true, false)
    assert.Equal(t, lock.IsSafe(), true)

    lock = mutex.New(true, true)
    assert.Equal(t, lock.IsSafe(), true)

    lock = mutex.New(true)
    assert.Equal(t, lock.IsSafe(), true)
}

func TestSafeMutex(t *testing.T) {
    safeLock := mutex.New(true)
    array := make([]any, 0, 10)

    go func() {
        safeLock.Lock()
        array = append(array, 1)
        time.Sleep(1000 * time.Millisecond)
        array = append(array, 1)
        safeLock.Unlock()
    }()
    go func() {
        time.Sleep(100 * time.Millisecond)
        safeLock.Lock()
        array = append(array, 1)
        time.Sleep(2000 * time.Millisecond)
        array = append(array, 1)
        safeLock.Unlock()
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

func TestUnsafeMutex(t *testing.T) {
    var (
        unsafeLock = mutex.New()
        array      = make([]any, 0, 10)
    )

    go func() {
        unsafeLock.Lock()
        array = append(array, 1)
        time.Sleep(1000 * time.Millisecond)
        array = append(array, 1)
        unsafeLock.Unlock()
    }()
    go func() {
        time.Sleep(100 * time.Millisecond)
        unsafeLock.Lock()
        array = append(array, 1)
        time.Sleep(2000 * time.Millisecond)
        array = append(array, 1)
        unsafeLock.Unlock()
    }()
    time.Sleep(500 * time.Millisecond)
    assert.Equal(t, len(array), 2)
    time.Sleep(1000 * time.Millisecond)
    assert.Equal(t, len(array), 3)
    time.Sleep(500 * time.Millisecond)
    assert.Equal(t, len(array), 3)
    time.Sleep(1000 * time.Millisecond)
    assert.Equal(t, len(array), 4)
}

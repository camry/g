package rwmutex

import (
    "sync"
)

// RWMutex is a sync.RWMutex with a switch for concurrent safe feature.
// If its attribute *sync.RWMutex is not nil, it means it's in concurrent safety usage.
// Its attribute *sync.RWMutex is nil in default, which makes this struct mush lightweight.
type RWMutex struct {
    // Underlying mutex.
    mutex *sync.RWMutex
}

// New creates and returns a new *RWMutex.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
func New(safe ...bool) *RWMutex {
    mu := Create(safe...)
    return &mu
}

// Create creates and returns a new RWMutex object.
// The parameter `safe` is used to specify whether using this mutex in concurrent safety,
// which is false in default.
func Create(safe ...bool) RWMutex {
    if len(safe) > 0 && safe[0] {
        return RWMutex{
            mutex: new(sync.RWMutex),
        }
    }
    return RWMutex{}
}

// IsSafe checks and returns whether current mutex is in concurrent-safe usage.
func (mu *RWMutex) IsSafe() bool {
    return mu.mutex != nil
}

// Lock locks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.
func (mu *RWMutex) Lock() {
    if mu.mutex != nil {
        mu.mutex.Lock()
    }
}

// Unlock unlocks mutex for writing.
// It does nothing if it is not in concurrent-safe usage.
func (mu *RWMutex) Unlock() {
    if mu.mutex != nil {
        mu.mutex.Unlock()
    }
}

// RLock locks mutex for reading.
// It does nothing if it is not in concurrent-safe usage.
func (mu *RWMutex) RLock() {
    if mu.mutex != nil {
        mu.mutex.RLock()
    }
}

// RUnlock unlocks mutex for reading.
// It does nothing if it is not in concurrent-safe usage.
func (mu *RWMutex) RUnlock() {
    if mu.mutex != nil {
        mu.mutex.RUnlock()
    }
}

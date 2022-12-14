package grand_test

import (
    "testing"

    "github.com/camry/g/gutil"
    "github.com/camry/g/gutil/grand"
    "github.com/stretchr/testify/assert"
)

func TestGRand_Hit(t *testing.T) {
    r := grand.NewRand(100000000)
    for i := 0; i < 100; i++ {
        assert.True(t, r.Hit(100, 100))
    }
    for i := 0; i < 100; i++ {
        assert.False(t, r.Hit(0, 100))
    }
    for i := 0; i < 10; i++ {
        assert.True(t, gutil.InArray(r.Hit(50, 100), []bool{true, false}))
    }
    for i := 0; i < 1000; i++ {
        assert.True(t, r.Hit(100, 100, 3))
    }
    for i := 0; i < 1000; i++ {
        assert.False(t, r.Hit(0, 100, 3))
    }
    for i := 0; i < 100; i++ {
        assert.True(t, gutil.InArray(r.Hit(50, 100, 3), []bool{true, false}))
    }
}

func TestGRand_HitProb(t *testing.T) {
    r := grand.NewRand(100000000)
    for i := 0; i < 100; i++ {
        assert.True(t, r.HitProb(1))
    }
    for i := 0; i < 100; i++ {
        assert.False(t, r.HitProb(0))
    }
    for i := 0; i < 100; i++ {
        assert.True(t, gutil.InArray(r.HitProb(0.5), []bool{true, false}))
    }
    for i := 0; i < 100; i++ {
        assert.True(t, r.HitProb(1, 3))
    }
    for i := 0; i < 100; i++ {
        assert.False(t, r.HitProb(0, 3))
    }
    for i := 0; i < 100; i++ {
        assert.True(t, gutil.InArray(r.HitProb(0.5, 3), []bool{true, false}))
    }
}

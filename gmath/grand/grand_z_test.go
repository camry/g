package grand_test

import (
    "testing"

    "github.com/samber/lo"
    "github.com/stretchr/testify/assert"

    "github.com/camry/g/gmath/grand"
)

func TestGRand_RangeInt(t *testing.T) {
    r := grand.NewRand(100000000)
    v1 := []int{1, 3, 3, 4, 9, 6, 2, 8, 9, 6}
    for i := 0; i < 10; i++ {
        assert.Equal(t, v1[i], r.RangeInt(0, 10))
    }
    v2 := []int{-17, -15, -15, -13, -12, -15, -13, -16, -10, -13, -11}
    for i := 0; i < 10; i++ {
        assert.Equal(t, v2[i], r.RangeInt(-20, -10))
    }
}

func TestGRand_Hit(t *testing.T) {
    r := grand.NewRand(100000000)
    for i := 0; i < 100; i++ {
        assert.True(t, r.Hit(100, 100))
    }
    for i := 0; i < 100; i++ {
        assert.False(t, r.Hit(0, 100))
    }
    for i := 0; i < 10; i++ {
        assert.True(t, lo.Contains([]bool{true, false}, r.Hit(50, 100)))
    }
    for i := 0; i < 1000; i++ {
        assert.True(t, r.Hit(100, 100, 3))
    }
    for i := 0; i < 1000; i++ {
        assert.False(t, r.Hit(0, 100, 3))
    }
    for i := 0; i < 100; i++ {
        assert.True(t, lo.Contains([]bool{true, false}, r.Hit(50, 100, 3)))
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
        assert.True(t, lo.Contains([]bool{true, false}, r.HitProb(0.5)))
    }
    for i := 0; i < 100; i++ {
        assert.True(t, r.HitProb(1, 3))
    }
    for i := 0; i < 100; i++ {
        assert.False(t, r.HitProb(0, 3))
    }
    for i := 0; i < 100; i++ {
        assert.True(t, lo.Contains([]bool{true, false}, r.HitProb(0.5, 3)))
    }
}

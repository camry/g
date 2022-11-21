package gutil_test

import (
    "testing"

    "github.com/camry/g/gutil"
    "github.com/stretchr/testify/assert"
)

func TestInArray(t *testing.T) {
    assert.True(t, gutil.InArray(1, []int{1, 2, 3, 4}), true)
    assert.False(t, gutil.InArray(5, []int{1, 2, 3, 4}), true)
    assert.True(t, gutil.InArray("cat", []string{"dog", "cat", "pig"}), true)
    assert.False(t, gutil.InArray("monkey", []string{"dog", "cat", "pig"}), true)
}

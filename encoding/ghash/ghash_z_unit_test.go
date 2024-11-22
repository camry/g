package ghash_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/encoding/ghash"
)

var (
    strBasic = []byte("This is the test string for hash.")
)

func Test_BKDR(t *testing.T) {
    x1 := uint32(200645773)
    j1 := ghash.BKDR(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(4214762819217104013)
    j2 := ghash.BKDR64(strBasic)
    assert.Equal(t, j2, x2)
}

func Test_SDBM(t *testing.T) {
    x1 := uint32(1069170245)
    j1 := ghash.SDBM(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(9881052176572890693)
    j2 := ghash.SDBM64(strBasic)
    assert.Equal(t, j2, x2)
}

func Test_RS(t *testing.T) {
    x1 := uint32(1944033799)
    j1 := ghash.RS(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(13439708950444349959)
    j2 := ghash.RS64(strBasic)
    assert.Equal(t, j2, x2)
}

func Test_JS(t *testing.T) {
    x1 := uint32(498688898)
    j1 := ghash.JS(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(13410163655098759877)
    j2 := ghash.JS64(strBasic)
    assert.Equal(t, j2, x2)
}

func Test_PJW(t *testing.T) {
    x1 := uint32(7244206)
    j1 := ghash.PJW(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(31150)
    j2 := ghash.PJW64(strBasic)
    assert.Equal(t, j2, x2)
}

func Test_ELF(t *testing.T) {
    x1 := uint32(7244206)
    j1 := ghash.ELF(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(31150)
    j2 := ghash.ELF64(strBasic)
    assert.Equal(t, j2, x2)
}

func Test_DJB(t *testing.T) {
    x1 := uint32(959862602)
    j1 := ghash.DJB(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(2519720351310960458)
    j2 := ghash.DJB64(strBasic)
    assert.Equal(t, j2, x2)
}

func Test_AP(t *testing.T) {
    x1 := uint32(3998202516)
    j1 := ghash.AP(strBasic)
    assert.Equal(t, j1, x1)

    x2 := uint64(2531023058543352243)
    j2 := ghash.AP64(strBasic)
    assert.Equal(t, j2, x2)
}

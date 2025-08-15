package gipv4_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gnet/gipv4"
)

func TestGetMac(t *testing.T) {
    mac, err := gipv4.GetMac()
    assert.Nil(t, err)
    assert.NotEqual(t, mac, "")
    // MAC addresses are typically 17 characters in length
    assert.Equal(t, len(mac), 17)
}

func TestGetMacArray(t *testing.T) {
    macs, err := gipv4.GetMacArray()
    assert.Nil(t, err)
    assert.Greater(t, len(macs), 0)
    for _, mac := range macs {
        // MAC addresses are typically 17 characters in length
        assert.Equal(t, len(mac), 17)
    }
}

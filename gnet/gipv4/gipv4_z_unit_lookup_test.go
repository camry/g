package gipv4_test

import (
    "testing"

    "github.com/samber/lo"
    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gnet/gipv4"
)

func TestGetHostByName(t *testing.T) {
    ip, err := gipv4.GetHostByName("localhost")
    assert.Nil(t, err)
    assert.Equal(t, ip, "127.0.0.1")
}

func TestGetHostsByName(t *testing.T) {
    ips, err := gipv4.GetHostsByName("localhost")
    assert.Nil(t, err)
    assert.True(t, lo.Contains(ips, "127.0.0.1"))
}

package gipv4_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gnet/gipv4"
)

const (
    ipv4             string = "192.168.1.1"
    longBigEndian    uint32 = 3232235777
    longLittleEndian uint32 = 16885952
)

func TestIpToLongBigEndian(t *testing.T) {
    var u = gipv4.IpToLongBigEndian(ipv4)
    assert.Equal(t, u, longBigEndian)

    var u2 = gipv4.Ip2long(ipv4)
    assert.Equal(t, u2, longBigEndian)
}

func TestLongToIpBigEndian(t *testing.T) {
    var s = gipv4.LongToIpBigEndian(longBigEndian)
    assert.Equal(t, s, ipv4)

    var s2 = gipv4.Long2ip(longBigEndian)
    assert.Equal(t, s2, ipv4)
}

func TestIpToLongLittleEndian(t *testing.T) {
    var u = gipv4.IpToLongLittleEndian(ipv4)
    assert.Equal(t, u, longLittleEndian)
}

func TestLongToIpLittleEndian(t *testing.T) {
    var s = gipv4.LongToIpLittleEndian(longLittleEndian)
    assert.Equal(t, s, ipv4)
}

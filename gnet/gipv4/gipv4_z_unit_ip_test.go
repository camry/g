package gipv4_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gnet/gipv4"
)

func TestGetIpArray(t *testing.T) {
    ips, err := gipv4.GetIpArray()
    assert.Nil(t, err)
    assert.Greater(t, len(ips), 0)
    for _, ip := range ips {
        assert.Equal(t, gipv4.Validate(ip), true)
    }
}

func TestMustGetIntranetIp(t *testing.T) {
    defer func() {
        if r := recover(); r != nil {
            t.Errorf("MustGetIntranetIp() panicked: %v", r)
        }
    }()
    ip := gipv4.MustGetIntranetIp()
    assert.Equal(t, gipv4.IsIntranet(ip), true)
}

func TestGetIntranetIp(t *testing.T) {
    ip, err := gipv4.GetIntranetIp()
    assert.Nil(t, err)
    assert.NotEqual(t, ip, "")
    assert.Equal(t, gipv4.IsIntranet(ip), true)
}

func TestGetIntranetIpArray(t *testing.T) {
    ips, err := gipv4.GetIntranetIpArray()
    assert.Nil(t, err)
    assert.Greater(t, len(ips), 0)
    for _, ip := range ips {
        assert.Equal(t, gipv4.IsIntranet(ip), true)
    }
}

func TestIsIntranet(t *testing.T) {
    tests := []struct {
        ip       string
        expected bool
    }{
        {"127.0.0.1", true},
        {"10.0.0.1", true},
        {"172.16.0.1", true},
        {"172.31.255.255", true},
        {"192.168.0.1", true},
        {"192.168.255.255", true},
        {"8.8.8.8", false},
        {"172.32.0.1", false},
        {"256.256.256.256", false},
    }

    for _, test := range tests {
        result := gipv4.IsIntranet(test.ip)
        assert.Equal(t, result, test.expected)
    }
}

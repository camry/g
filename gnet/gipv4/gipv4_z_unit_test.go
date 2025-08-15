package gipv4_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gnet/gipv4"
)

func TestValidate(t *testing.T) {
    tests := []struct {
        ip       string
        expected bool
    }{
        {"192.168.1.1", true},
        {"255.255.255.255", true},
        {"0.0.0.0", true},
        {"256.256.256.256", false},
        {"192.168.1", false},
        {"abc.def.ghi.jkl", false},
        {"19216811", false},
        {"abcdefghijkl", false},
    }

    for _, test := range tests {
        result := gipv4.Validate(test.ip)
        assert.Equal(t, result, test.expected)
    }
}

func TestParseAddress(t *testing.T) {
    tests := []struct {
        address      string
        expectedIP   string
        expectedPort int
    }{
        {"192.168.1.1:80", "192.168.1.1", 80},
        {"10.0.0.1:8080", "10.0.0.1", 8080},
        {"127.0.0.1:65535", "127.0.0.1", 65535},
        {"invalid:address", "", 0},
        {"192.168.1.1", "", 0},
        {"19216811", "", 0},
    }

    for _, test := range tests {
        ip, port := gipv4.ParseAddress(test.address)
        assert.Equal(t, ip, test.expectedIP)
        assert.Equal(t, port, test.expectedPort)
    }
}

func TestGetSegment(t *testing.T) {
    tests := []struct {
        ip       string
        expected string
    }{
        {"192.168.2.102", "192.168.2"},
        {"10.0.0.1", "10.0.0"},
        {"255.255.255.255", "255.255.255"},
        {"invalid.ip.address", ""},
        {"123", ""},
        {"192.168.2.102.123", ""},
    }

    for _, test := range tests {
        result := gipv4.GetSegment(test.ip)
        assert.Equal(t, result, test.expected)
    }
}

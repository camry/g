package gudp_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/gnet/gudp"
)

func TestGetFreePort(t *testing.T) {
    _, err := gudp.GetFreePort()
    if err != nil {
        t.Error(err)
    }
}

func TestGetFreePorts(t *testing.T) {
    ports, err := gudp.GetFreePorts(2)
    if err != nil {
        t.Error(err)
    }
    assert.Equal(t, len(ports), 2)
}

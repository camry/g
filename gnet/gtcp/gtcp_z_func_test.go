package gtcp_test

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/gnet/gtcp"
)

func TestGetFreePort(t *testing.T) {
    _, err := gtcp.GetFreePort()
    if err != nil {
        t.Error(err)
    }
}

func TestGetFreePorts(t *testing.T) {
    ports, err := gtcp.GetFreePorts(2)
    if err != nil {
        t.Error(err)
    }
    assert.Equal(t, len(ports), 2)
}

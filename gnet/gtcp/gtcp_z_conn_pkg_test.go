package gtcp_test

import (
    "context"
    "github.com/stretchr/testify/assert"
    "strconv"
    "testing"
    "time"

    "github.com/camry/g/gnet/gtcp"
)

func TestPackageBasic(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.ReceivePkg()
            if err != nil {
                break
            }
            conn.SendPkg(data)
        }
    }, nil)
    go s.Run(ctx)
    defer s.Close(ctx)
    time.Sleep(100 * time.Millisecond)

    // SendPkg
    conn1, err1 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err1)
    defer conn1.Close()
    for i := 0; i < 100; i++ {
        err1 := conn1.SendPkg([]byte(strconv.Itoa(i)))
        assert.Nil(t, err1)
    }
    for i := 0; i < 100; i++ {
        err1 := conn1.SendPkgWithTimeout([]byte(strconv.Itoa(i)), time.Second)
        assert.Nil(t, err1)
    }

    // SendPkg with big data - failure.
    conn2, err2 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err2)
    defer conn2.Close()
    data2 := make([]byte, 65536)
    err2 = conn2.SendPkg(data2)
    assert.NotNil(t, err2)

    // SendReceivePkg
    conn3, err3 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err3)
    defer conn3.Close()
    for i := 100; i < 200; i++ {
        data3 := []byte(strconv.Itoa(i))
        result3, err3 := conn3.SendReceivePkg(data3)
        assert.Nil(t, err3)
        assert.Equal(t, result3, data3)
    }
    // SendReceivePkgWithTimeout
    for i := 100; i < 200; i++ {
        data3 := []byte(strconv.Itoa(i))
        result3, err3 := conn3.SendReceivePkgWithTimeout(data3, time.Second)
        assert.Nil(t, err3)
        assert.Equal(t, result3, data3)
    }

    // SendReceivePkg with big data - failure.
    conn4, err4 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err4)
    defer conn4.Close()
    data4 := make([]byte, 65536)
    result4, err4 := conn4.SendReceivePkg(data4)
    assert.NotNil(t, err4)
    assert.Nil(t, result4)

    // SendReceivePkg with big data - success.
    conn5, err5 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err5)
    defer conn5.Close()
    data5 := make([]byte, 65500)
    data5[100] = byte(65)
    data5[65400] = byte(85)
    result5, err5 := conn5.SendReceivePkg(data5)
    assert.Nil(t, err5)
    assert.Equal(t, result5, data5)
}

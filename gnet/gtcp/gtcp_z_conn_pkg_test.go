package gtcp_test

import (
    "context"
    "strconv"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"

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
    })
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

func TestPackageBasicHeaderSize1(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.ReceivePkg(gtcp.PkgOption{HeaderSize: 1})
            if err != nil {
                break
            }
            conn.SendPkg(data)
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)
    time.Sleep(100 * time.Millisecond)

    // SendReceivePkg with empty data.
    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    defer conn.Close()
    data := make([]byte, 0)
    result, err := conn.SendReceivePkg(data)
    assert.Nil(t, err)
    assert.Nil(t, result)
}

func TestPackageTimeout(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.ReceivePkg()
            if err != nil {
                break
            }
            time.Sleep(time.Second)
            assert.Nil(t, conn.SendPkg(data))
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)
    time.Sleep(100 * time.Millisecond)

    // SendReceivePkgWithTimeout
    conn1, err1 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err1)
    defer conn1.Close()
    data1 := []byte("10000")
    result1, err1 := conn1.SendReceivePkgWithTimeout(data1, time.Millisecond*500)
    assert.NotNil(t, err1)
    assert.Nil(t, result1)

    // SendReceivePkgWithTimeout
    conn2, err2 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err2)
    defer conn2.Close()
    data2 := []byte("10000")
    result2, err2 := conn2.SendReceivePkgWithTimeout(data2, time.Second*2)
    assert.Nil(t, err2)
    assert.Equal(t, result2, data2)

    // SendReceivePkgWithTimeout
    conn3, err3 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err3)
    defer conn3.Close()
    data3 := []byte("10000")
    result3, err3 := conn3.SendReceivePkgWithTimeout(data3, time.Second*2, gtcp.PkgOption{HeaderSize: 5})
    assert.NotNil(t, err3)
    assert.Nil(t, result3)
}

func TestPackageOption(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        option := gtcp.PkgOption{HeaderSize: 1}
        for {
            data, err := conn.ReceivePkg(option)
            if err != nil {
                break
            }
            assert.Nil(t, conn.SendPkg(data, option))
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)
    time.Sleep(100 * time.Millisecond)

    // SendReceivePkg with big data - failure.
    conn1, err1 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err1)
    defer conn1.Close()
    data1 := make([]byte, 0xFF+1)
    result1, err1 := conn1.SendReceivePkg(data1, gtcp.PkgOption{HeaderSize: 1})
    assert.NotNil(t, err1)
    assert.Nil(t, result1)

    // SendReceivePkg with big data - success.
    conn2, err2 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err2)
    defer conn2.Close()
    data2 := make([]byte, 0xFF)
    data2[100] = byte(65)
    data2[200] = byte(85)
    result2, err2 := conn2.SendReceivePkg(data2, gtcp.PkgOption{HeaderSize: 1})
    assert.Nil(t, err2)
    assert.Equal(t, result2, data2)
}

func TestPackageOptionHeadSize3(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        option := gtcp.PkgOption{HeaderSize: 3}
        for {
            data, err := conn.ReceivePkg(option)
            if err != nil {
                break
            }
            assert.Nil(t, conn.SendPkg(data, option))
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)
    time.Sleep(100 * time.Millisecond)

    // SendReceivePkg
    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    defer conn.Close()
    data := make([]byte, 0xFF)
    data[100] = byte(65)
    data[200] = byte(85)
    result, err := conn.SendReceivePkg(data, gtcp.PkgOption{HeaderSize: 3})
    assert.Nil(t, err)
    assert.Equal(t, result, data)
}

func Test_Package_Option_HeadSize4(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        option := gtcp.PkgOption{HeaderSize: 4}
        for {
            data, err := conn.ReceivePkg(option)
            if err != nil {
                break
            }
            assert.Nil(t, conn.SendPkg(data, option))
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)
    time.Sleep(100 * time.Millisecond)
    // SendReceivePkg with big data - failure.
    conn1, err1 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err1)
    defer conn1.Close()
    data1 := make([]byte, 0xFFFF+1)
    _, err1 = conn1.SendReceivePkg(data1, gtcp.PkgOption{HeaderSize: 4})
    assert.Nil(t, err1)

    // SendReceivePkg with big data - success.
    conn2, err2 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err2)
    defer conn2.Close()
    data2 := make([]byte, 0xFF)
    data2[100] = byte(65)
    data2[200] = byte(85)
    result2, err2 := conn2.SendReceivePkg(data2, gtcp.PkgOption{HeaderSize: 4})
    assert.Nil(t, err2)
    assert.Equal(t, result2, data2)

    // pkgOption.HeaderSize oversize
    conn3, err3 := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err3)
    defer conn3.Close()
    data3 := make([]byte, 0xFF)
    data3[100] = byte(65)
    data3[200] = byte(85)
    _, err3 = conn3.SendReceivePkg(data3, gtcp.PkgOption{HeaderSize: 5})
    assert.NotNil(t, err3)
}

func TestConnReceivePkgError(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        option := gtcp.PkgOption{HeaderSize: 5}
        for {
            _, err := conn.ReceivePkg(option)
            if err != nil {
                break
            }
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)

    time.Sleep(100 * time.Millisecond)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    defer conn.Close()
    data := make([]byte, 65536)
    result, err := conn.SendReceivePkg(data)
    assert.NotNil(t, err)
    assert.Nil(t, result)
}

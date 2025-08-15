package gtcp_test

import (
    "context"
    "fmt"
    "strings"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/gnet/gtcp"
)

var (
    simpleTimeout = time.Millisecond * 100
    sendData      = []byte("hello")
    invalidAddr   = "127.0.0.1:99999"
)

func startTCPServer(addr string) *gtcp.Server {
    ctx := context.Background()
    s := gtcp.NewServer(addr, func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.Recv(-1)
            if err != nil {
                break
            }
            conn.Send(data)
        }
    })
    go s.Run(ctx)
    time.Sleep(simpleTimeout)
    return s
}

func startTCPPkgServer(addr string) *gtcp.Server {
    ctx := context.Background()
    s := gtcp.NewServer(addr, func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.RecvPkg()
            if err != nil {
                break
            }
            conn.SendPkg(data)
        }
    })
    go s.Run(ctx)
    time.Sleep(simpleTimeout)
    return s
}

func TestConnGetFreePorts(t *testing.T) {
    ports, _ := gtcp.GetFreePorts(2)
    assert.Greater(t, ports[0], 0)
    assert.Greater(t, ports[1], 0)

    startTCPServer(fmt.Sprintf("%s:%d", "127.0.0.1", ports[0]))

    conn, err := gtcp.NewConn(fmt.Sprintf("127.0.0.1:%d", ports[0]))
    assert.Nil(t, err)
    defer conn.Close()
    result, err := conn.SendRecv(sendData, -1)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)

    conn1, err1 := gtcp.NewConn(fmt.Sprintf("127.0.0.1:%d", 80))
    assert.NotNil(t, err1)
    assert.Nil(t, conn1)
}

func TestConnMustGetFreePort(t *testing.T) {
    port := gtcp.MustGetFreePort()
    addr := fmt.Sprintf("%s:%d", "127.0.0.1", port)
    startTCPServer(addr)

    result, err := gtcp.SendRecv(addr, sendData, -1)
    assert.Nil(t, err)
    assert.Equal(t, sendData, result)
}

func TestNewConn(t *testing.T) {
    addr := gtcp.FreePortAddress

    conn, err := gtcp.NewConn(addr, simpleTimeout)
    assert.Nil(t, conn)
    assert.NotNil(t, err)

    s := startTCPServer(gtcp.FreePortAddress)

    conn1, err1 := gtcp.NewConn(s.GetListenedAddress(), simpleTimeout)
    assert.Nil(t, err1)
    assert.NotEqual(t, conn1, nil)
    defer conn1.Close()
    result1, err1 := conn1.SendRecv(sendData, -1)
    assert.Nil(t, err1)
    assert.Equal(t, result1, sendData)
}

func TestConn_Send(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    err = conn.Send(sendData, gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    result, err := conn.Recv(-1)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestConn_SendWithTimeout(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    err = conn.SendWithTimeout(sendData, time.Second, gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    result, err := conn.Recv(-1)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestConn_SendRecv(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    result, err := conn.SendRecv(sendData, -1, gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestConn_SendRecvWithTimeout(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    result, err := conn.SendRecvWithTimeout(sendData, -1, time.Second, gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestConn_RecvWithTimeout(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    conn.Send(sendData)
    result, err := conn.RecvWithTimeout(-1, time.Second, gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestConn_RecvLine(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    data := []byte("hello\n")
    conn.Send(data)
    result, err := conn.RecvLine(gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    splitData := strings.Split(string(data), "\n")
    assert.Equal(t, string(result), splitData[0])
}

func TestConn_RecvTill(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    conn.Send(sendData)
    result, err := conn.RecvTill([]byte("hello"), gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestConn_SetDeadline(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    conn.SetDeadline(time.Time{})
    err = conn.Send(sendData, gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    result, err := conn.Recv(-1)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestConn_SetRecvBufferWait(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    conn, err := gtcp.NewConn(s.GetListenedAddress())
    assert.Nil(t, err)
    assert.NotNil(t, conn)
    conn.SetBufferWaitRecv(time.Millisecond * 100)
    err = conn.Send(sendData, gtcp.Retry{Count: 1})
    assert.Nil(t, err)
    result, err := conn.Recv(-1)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestSend(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    err1 := gtcp.Send(invalidAddr, sendData, gtcp.Retry{Count: 1})
    assert.NotNil(t, err1)

    err2 := gtcp.Send(s.GetListenedAddress(), sendData, gtcp.Retry{Count: 1})
    assert.Nil(t, err2)
}

func TestSendRecv(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    result1, err1 := gtcp.SendRecv(invalidAddr, sendData, -1)
    assert.NotNil(t, err1)
    assert.Nil(t, result1)

    result2, err2 := gtcp.SendRecv(s.GetListenedAddress(), sendData, -1)
    assert.Nil(t, err2)
    assert.Equal(t, result2, sendData)
}

func TestSendWithTimeout(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    err := gtcp.SendWithTimeout(invalidAddr, sendData, time.Millisecond*500)
    assert.NotNil(t, err)
    err = gtcp.SendWithTimeout(s.GetListenedAddress(), sendData, time.Millisecond*500)
    assert.Nil(t, err)
}

func TestSendRecvWithTimeout(t *testing.T) {
    s := startTCPServer(gtcp.FreePortAddress)

    result, err := gtcp.SendRecvWithTimeout(invalidAddr, sendData, -1, time.Millisecond*500)
    assert.Nil(t, result)
    assert.NotNil(t, err)
    result, err = gtcp.SendRecvWithTimeout(s.GetListenedAddress(), sendData, -1, time.Millisecond*500)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestSendPkg(t *testing.T) {
    s := startTCPPkgServer(gtcp.FreePortAddress)

    err1 := gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err1)
    err1 = gtcp.SendPkg(invalidAddr, sendData)
    assert.NotNil(t, err1)

    err2 := gtcp.SendPkg(s.GetListenedAddress(), sendData, gtcp.PkgOption{Retry: gtcp.Retry{Count: 3}})
    assert.Nil(t, err2)
    err2 = gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err2)
}

func TestSendRecvPkg(t *testing.T) {
    s := startTCPPkgServer(gtcp.FreePortAddress)

    err1 := gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err1)
    _, err1 = gtcp.SendRecvPkg(invalidAddr, sendData)
    assert.NotNil(t, err1)

    err2 := gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err2)
    result, err2 := gtcp.SendRecvPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err2)
    assert.Equal(t, result, sendData)
}

func TestSendPkgWithTimeout(t *testing.T) {
    s := startTCPPkgServer(gtcp.FreePortAddress)

    err1 := gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err1)
    err1 = gtcp.SendPkgWithTimeout(invalidAddr, sendData, time.Second)
    assert.NotNil(t, err1)

    err2 := gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err2)
    err2 = gtcp.SendPkgWithTimeout(s.GetListenedAddress(), sendData, time.Second)
    assert.Nil(t, err2)
}

func TestSendRecvPkgWithTimeout(t *testing.T) {
    s := startTCPPkgServer(gtcp.FreePortAddress)

    err1 := gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err1)
    _, err1 = gtcp.SendRecvPkgWithTimeout(invalidAddr, sendData, time.Second)
    assert.NotNil(t, err1)

    err2 := gtcp.SendPkg(s.GetListenedAddress(), sendData)
    assert.Nil(t, err2)
    result, err2 := gtcp.SendRecvPkgWithTimeout(s.GetListenedAddress(), sendData, time.Second)
    assert.Nil(t, err2)
    assert.Equal(t, result, sendData)
}

func TestNewServer(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.Recv(-1)
            if err != nil {
                break
            }
            conn.Send(data)
        }
    })
    defer s.Close(ctx)
    go s.Run(ctx)

    time.Sleep(simpleTimeout)

    result, err := gtcp.SendRecv(s.GetListenedAddress(), sendData, -1)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)
}

func TestServer_Run(t *testing.T) {
    ctx := context.Background()
    s := gtcp.NewServer(gtcp.FreePortAddress, func(conn *gtcp.Conn) {
        defer conn.Close()
        for {
            data, err := conn.Recv(-1)
            if err != nil {
                break
            }
            conn.Send(data)
        }
    })
    defer s.Close(ctx)
    go s.Run(ctx)

    time.Sleep(simpleTimeout)

    result, err := gtcp.SendRecv(s.GetListenedAddress(), sendData, -1)
    assert.Nil(t, err)
    assert.Equal(t, result, sendData)

    s1 := gtcp.NewServer(gtcp.FreePortAddress, nil)
    defer s1.Close(ctx)
    go func() {
        err := s1.Run(ctx)
        assert.NotNil(t, err)
    }()
}

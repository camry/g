// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gudp_test

import (
    "context"
    "fmt"
    "io"
    "strconv"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"

    "github.com/camry/g/v2/glog"
    "github.com/camry/g/v2/gnet/gudp"
)

var (
    simpleTimeout = time.Millisecond * 100
    sendData      = []byte("hello")
)

func startUDPServer(addr string) *gudp.Server {
    s := gudp.NewServer(addr, func(conn *gudp.ServerConn) {
        defer conn.Close()
        for {
            data, remote, err := conn.Recv(-1)
            if err != nil {
                if err != io.EOF {
                    glog.Error(context.TODO(), err)
                }
                break
            }
            if err = conn.Send(data, remote); err != nil {
                glog.Error(context.TODO(), err)
            }
        }
    })
    go s.Run(context.Background())
    time.Sleep(simpleTimeout)
    return s
}

func Test_Basic(t *testing.T) {
    var ctx = context.TODO()
    s := gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.ServerConn) {
        defer conn.Close()
        for {
            data, remote, err := conn.Recv(-1)
            if len(data) > 0 {
                if err = conn.Send(append([]byte("> "), data...), remote); err != nil {
                    glog.Error(ctx, err)
                }
            }
            if err != nil {
                break
            }
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)

    time.Sleep(100 * time.Millisecond)

    // gudp.Conn.Send
    for i := 0; i < 100; i++ {
        conn, err := gudp.NewClientConn(s.GetListenedAddress())
        assert.Nil(t, err)
        assert.Equal(t, conn.Send([]byte(strconv.Itoa(i))), nil)
        assert.NotEqual(t, conn.RemoteAddr(), nil)
        result, _, err := conn.Recv(-1)
        assert.Nil(t, err)
        assert.NotEqual(t, conn.RemoteAddr(), nil)
        assert.Equal(t, string(result), fmt.Sprintf(`> %d`, i))
        conn.Close()
    }
    // gudp.Conn.SendRecv
    for i := 0; i < 100; i++ {
        conn, err := gudp.NewClientConn(s.GetListenedAddress())
        assert.Nil(t, err)
        result, err := conn.SendRecv([]byte(strconv.Itoa(i)), -1)
        assert.Nil(t, err)
        assert.Equal(t, string(result), fmt.Sprintf(`> %d`, i))
        conn.Close()
    }

    // gudp.Send
    for i := 0; i < 100; i++ {
        err := gudp.Send(s.GetListenedAddress(), []byte(strconv.Itoa(i)))
        assert.Nil(t, err)
    }
}

func Test_Buffer(t *testing.T) {
    var ctx = context.TODO()
    s := gudp.NewServer(gudp.FreePortAddress, func(conn *gudp.ServerConn) {
        defer conn.Close()
        for {
            data, remote, err := conn.Recv(-1)
            if len(data) > 0 {
                if err = conn.Send(data, remote); err != nil {
                    glog.Error(ctx, err)
                }
            }
            if err != nil {
                break
            }
        }
    })
    go s.Run(ctx)
    defer s.Close(ctx)
    time.Sleep(100 * time.Millisecond)

    result, err := gudp.SendRecv(s.GetListenedAddress(), []byte("123"), -1)
    assert.Nil(t, err)
    assert.Equal(t, string(result), "123")

    result, err = gudp.SendRecv(s.GetListenedAddress(), []byte("456"), -1)
    assert.Nil(t, err)
    assert.Equal(t, string(result), "456")
}

func Test_NewConn(t *testing.T) {
    s := startUDPServer(gudp.FreePortAddress)

    conn, err := gudp.NewClientConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", gudp.MustGetFreePort()))
    assert.Nil(t, err)
    conn.SetDeadline(time.Now().Add(time.Second))
    assert.Equal(t, conn.Send(sendData), nil)
    conn.Close()

    conn, err = gudp.NewClientConn(s.GetListenedAddress(), fmt.Sprintf("127.0.0.1:%d", 99999))
    assert.Nil(t, conn)
    assert.NotEqual(t, err, nil)

    conn, err = gudp.NewClientConn(fmt.Sprintf("127.0.0.1:%d", 99999))
    assert.Nil(t, conn)
    assert.NotEqual(t, err, nil)
}

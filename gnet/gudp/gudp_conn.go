package gudp

import (
    "io"
    "net"
    "time"

    "github.com/camry/g/v2/gerrors/gerror"
)

// localConn 对 UDP 连接提供公共操作。
type localConn struct {
    *net.UDPConn           // 底层UDP连接。
    deadlineRecv time.Time // 读取超时点。
    deadlineSend time.Time // 写入超时点。
}

const (
    defaultRetryInterval  = 100 * time.Millisecond // 默认重试间隔。
    defaultReadBufferSize = 1024                   // (Byte)默认读取缓冲区大小。
)

// Retry 重试选项。
// TODO replace with standalone retry package.
type Retry struct {
    Count    int           // Max retry count.
    Interval time.Duration // Retry interval.
}

// Recv 接收并返回来自远程地址的数据。
// buffer 参数用于定制接收缓冲区的大小。
// 如果 `buffer` <= 0，它使用默认的缓冲区大小，即1024字节。
//
// 在UDP协议中有包边界，如果指定缓冲区大小足够大，我们可以接收一个完整的包。
// 非常注意，我们应该一次收到完整的包裹，否则剩余的包数据将被丢弃。
func (c *localConn) Recv(buffer int, retry ...Retry) ([]byte, *net.UDPAddr, error) {
    var (
        err        error        // 读取错误。
        size       int          // 读取大小。
        data       []byte       // 缓存数据。
        remoteAddr *net.UDPAddr // 要读取的当前远程地址。
    )
    if buffer > 0 {
        data = make([]byte, buffer)
    } else {
        data = make([]byte, defaultReadBufferSize)
    }
    for {
        size, remoteAddr, err = c.ReadFromUDP(data)
        if err != nil {
            // 连接关闭。
            if err == io.EOF {
                break
            }
            if len(retry) > 0 {
                // 即使重试也失败了。
                if retry[0].Count == 0 {
                    break
                }
                retry[0].Count--
                if retry[0].Interval == 0 {
                    retry[0].Interval = defaultRetryInterval
                }
                time.Sleep(retry[0].Interval)
                continue
            }
            err = gerror.Wrap(err, `ReadFromUDP failed`)
            break
        }
        break
    }
    return data[:size], remoteAddr, err
}

// SetDeadline 设置当前连接的读取和写入的截止时间点。
func (c *localConn) SetDeadline(t time.Time) (err error) {
    if err = c.UDPConn.SetDeadline(t); err == nil {
        c.deadlineRecv = t
        c.deadlineSend = t
    } else {
        err = gerror.Wrapf(err, `SetDeadline for connection failed with "%s"`, t)
    }
    return err
}

// SetDeadlineRecv 设置当前连接的读取截止时间点。
func (c *localConn) SetDeadlineRecv(t time.Time) (err error) {
    if err = c.SetReadDeadline(t); err == nil {
        c.deadlineRecv = t
    } else {
        err = gerror.Wrapf(err, `SetDeadlineRecv for connection failed with "%s"`, t)
    }
    return err
}

// SetDeadlineSend 设置当前连接的写入截止时间点。
func (c *localConn) SetDeadlineSend(t time.Time) (err error) {
    if err = c.SetWriteDeadline(t); err == nil {
        c.deadlineSend = t
    } else {
        err = gerror.Wrapf(err, `SetDeadlineSend for connection failed with "%s"`, t)
    }
    return err
}

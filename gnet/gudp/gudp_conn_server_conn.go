package gudp

import (
    "io"
    "net"
    "time"

    "github.com/camry/g/gerrors/gerror"
)

// ServerConn 服务端连接。
type ServerConn struct {
    *localConn
}

// NewServerConn 创建一个监听 `localAddress` 的UDP连接。
func NewServerConn(listenedConn *net.UDPConn) *ServerConn {
    return &ServerConn{
        localConn: &localConn{
            UDPConn: listenedConn,
        },
    }
}

// Send 将数据写入远程地址。
func (c *ServerConn) Send(data []byte, remoteAddr *net.UDPAddr, retry ...Retry) (err error) {
    for {
        _, err = c.WriteToUDP(data, remoteAddr)
        if err == nil {
            return nil
        }
        // 连接关闭。
        if err == io.EOF {
            return err
        }
        // 重试后仍然失败。
        if len(retry) == 0 || retry[0].Count == 0 {
            return gerror.Wrap(err, `Write data failed`)
        }
        if len(retry) > 0 {
            retry[0].Count--
            if retry[0].Interval == 0 {
                retry[0].Interval = defaultRetryInterval
            }
            time.Sleep(retry[0].Interval)
            continue
        }
        return err
    }
}
